package channels

import (
	"github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
	typeserr "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
)

var _ IChannel = new(Channel)

type IChannel interface {
	UpdateClient() error
	Relay() error
	IsNotRelay() bool
	Context() *domain.Context
	UpdateClientFrequency() uint64
}

type Channel struct {
	source repostitory.IChain
	dest   repostitory.IChain

	height  uint64
	context *domain.Context

	logger *log.Logger
}

func NewChannel(source repostitory.IChain, dest repostitory.IChain, height uint64) IChannel {

	return &Channel{
		logger:  log.New(),
		source:  source,
		dest:    dest,
		height:  height,
		context: domain.NewContext(height, source.ChainName()),
	}
}

func (channel *Channel) UpdateClientFrequency() uint64 {
	return channel.source.UpdateClientFrequency()
}

func (channel *Channel) UpdateClient() error {
	// 1. get light client state from dest chain
	logger := channel.logger.WithFields(log.Fields{
		"source_chain": channel.source.ChainName(),
		"dest_chain":   channel.dest.ChainName(),
		"option":       "cron_update_client",
	})
	clientState, err := channel.dest.GetLightClientState(channel.source.ChainName())
	if err != nil {
		logger.WithFields(log.Fields{
			"err_msg": err.Error(),
		}).Error("failed to get light client state")
		return typeserr.ErrGetLightClientState
	}
	// 2. get source chain updated latest height from dest chain
	heightObj := clientState.GetLatestHeight()
	height := heightObj.GetRevisionHeight()

	nextHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		logger.WithFields(log.Fields{
			"err_msg": err.Error(),
		}).Error("failed to get block header")
		return typeserr.ErrGetBlockHeader
	}
	return channel.updateClient(height, nextHeight)
}

func (channel *Channel) updateClient(trustedHeight, latestHeight uint64) error {
	// 3. get nextHeight block header from source chain
	logger := channel.logger.WithFields(log.Fields{
		"trusted_height": trustedHeight,
		"latest_height":  latestHeight,
		"source_chain":   channel.source.ChainName(),
		"dest_chain":     channel.dest.ChainName(),
		"option":         "update_client",
	})
	var header tibctypes.Header
	var err error
	switch channel.source.ChainType() {
	case constant.Tendermint:
		req := &repostitory.GetBlockHeaderReq{
			LatestHeight:  latestHeight,
			TrustedHeight: trustedHeight,
		}
		header, err = channel.source.GetBlockHeader(req)
		if err != nil {
			logger.Error("failed to get block header")
			return typeserr.ErrGetBlockHeader
		}
	}

	// 4. update client to dest chain
	hash, err := channel.dest.UpdateClient(header, channel.source.ChainName())
	if err != nil {
		channel.logger.Error("failed to update client")
		return typeserr.ErrUpdateClient
	}
	logger.WithFields(log.Fields{"dest_hash": hash}).Info()

	return nil
}

func (channel *Channel) Relay() error {

	logger := channel.logger.WithFields(log.Fields{
		"source_height": channel.context.Height(),
		"source_chain":  channel.source.ChainName(),
		"dest_chain":    channel.dest.ChainName(),
		"option":        "relay",
	})
	latestHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		channel.logger.Error("failed to get latest height")
		return typeserr.ErrGetLatestHeight
	}

	if latestHeight <= channel.context.Height() {
		channel.logger.Info("the current height cannot be relayed yet")
		return typeserr.ErrNotProduced
	}

	// 1. get packets from source chain
	packets, err := channel.source.GetPackets(channel.context.Height())
	if err != nil {
		logger.WithFields(log.Fields{
			"err_msg": err.Error(),
		}).Error("failed to get packets")
		return typeserr.ErrGetPackets
	}

	clientState, err := channel.dest.GetLightClientState(channel.source.ChainName())
	if err != nil {
		logger.WithFields(log.Fields{
			"err_msg": err.Error(),
		}).Error("failed to get light client state")
		return typeserr.ErrGetLightClientState
	}

	// 2.  Process biz packets
	var recvPackets types.Msgs
	for _, pack := range packets.BizPackets {

		// If packet.sourceChain == channel.dest.ChainName(),
		// Indicates that the current packet is sent by dest.
		// So data packets should not be relayed back
		if pack.SourceChain == channel.dest.ChainName() {
			continue
		}

		// 2.1 get commitments packets from source chain
		// The source and dest in the packet must be used here
		// commitment path is determined
		commitmentsPacketResp, err := channel.source.GetCommitmentsPacket(
			pack.SourceChain, pack.DestinationChain, pack.Sequence)
		if err != nil {
			logger.WithFields(log.Fields{
				"err_msg": err.Error(),
			}).Error("failed to get commitment packet")
			return typeserr.ErrGetCommitmentPacket
		}

		// 2.2 get receipt packet from dest chain
		// The source and dest in the packet must be used here
		// commitment path is determined
		receiptPacketResp, err := channel.dest.GetReceiptPacket(pack.SourceChain, pack.DestinationChain, pack.Sequence)
		if err != nil {
			logger.WithFields(log.Fields{
				"err_msg": err.Error(),
			}).Error("failed to get receipt packet")
			return typeserr.ErrGetReceiptPacket
		}
		// if receipt exist, skip
		if receiptPacketResp.Received {
			logger.Info("receipt exist, skip cur height ")
			continue
		}

		proof, err := channel.source.GetProof(
			pack.SourceChain, pack.DestinationChain, pack.Sequence,
			latestHeight, repostitory.CommitmentPoof)
		if err != nil {
			logger.Error("failed to get proof")
			return typeserr.ErrGetProof
		}

		recvPacket := &packet.MsgRecvPacket{
			Packet:          pack,
			ProofCommitment: proof,
			ProofHeight: client.Height{
				RevisionNumber: commitmentsPacketResp.ProofHeight.RevisionNumber,
				RevisionHeight: latestHeight,
			},
		}
		recvPackets = append(recvPackets, recvPacket)
	}

	//3. Process ack packets
	for _, pack := range packets.AckPackets {
		// query proof

		// If packet.DestinationChain == channel.source.ChainName(),
		// Indicates that the current packet is sent by dest.
		// So data packets should not be relayed back

		_, err := channel.dest.GetCommitmentsPacket(
			pack.Packet.SourceChain, pack.Packet.DestinationChain, pack.Packet.Sequence)
		if err != nil {
			logger.Info("the current packet has been confirmed")
			continue
		}

		proof, err := channel.source.GetProof(
			pack.Packet.SourceChain,
			pack.Packet.DestinationChain,
			pack.Packet.Sequence, latestHeight, repostitory.AckProof)
		if err != nil {
			logger.WithFields(log.Fields{
				"err_msg": err.Error(),
			}).Error("failed to get proof")
			return typeserr.ErrGetProof
		}
		recvPacket := &packet.MsgAcknowledgement{
			Packet:          pack.Packet,
			Acknowledgement: pack.Acknowledgement,
			ProofAcked:      proof,
			ProofHeight: client.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: latestHeight,
			},
		}
		recvPackets = append(recvPackets, recvPacket)
	}

	for _, pack := range packets.CleanPackets {
		recvPacket := &packet.MsgCleanPacket{
			CleanPacket: pack,
		}
		recvPackets = append(recvPackets, recvPacket)
	}

	if (len(packets.CleanPackets) == 0 && len(packets.AckPackets) == 0 && len(packets.BizPackets) == 0) || len(recvPackets) == 0 {
		logger.Info("there are no packets to be relayed at the current altitude")
		channel.Context().IncrHeight()
		return nil
	}

	// boastCommit tx
	err = channel.updateClient(clientState.GetLatestHeight().GetRevisionHeight(), latestHeight)
	if err != nil {
		return typeserr.ErrUpdateClient
	}
	resultTx, err := channel.dest.RecvPackets(recvPackets)
	if err != nil {
		logger.WithFields(log.Fields{
			"err_msg": err.Error(),
		}).Error("failed to recv packet")
		return typeserr.ErrRecvPacket
	}
	logger.WithFields(log.Fields{
		"tx_height":  resultTx.Height,
		"tx_hash":    resultTx.Hash,
		"gas_wanted": resultTx.GasWanted,
		"gas_used":   resultTx.GasUsed,
	}).Info("success")
	return nil
}

func (channel *Channel) IsNotRelay() bool {
	curHeight := channel.context.Height()
	latestHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		return false
	}

	if curHeight < latestHeight {
		return true
	}

	return false
}

func (channel *Channel) Context() *domain.Context {
	return channel.context
}
