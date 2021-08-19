package channels

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
	typeserr "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	"github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
	log "github.com/sirupsen/logrus"
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
	signer  string
	context *domain.Context

	logger log.Logger
}

func NewChannel(source repostitory.IChain, dest repostitory.IChain, height uint64) IChannel {

	return &Channel{
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
	clientState, err := channel.dest.GetLightClientState(channel.source.ChainName())
	if err != nil {
		channel.logger.Error("failed to get light client state")
		return typeserr.ErrGetLightClientState
	}

	// 2. get source chain updated latest height from dest chain
	heightObj := clientState.GetLatestHeight()
	height := heightObj.GetRevisionHeight()

	logger := channel.logger.WithFields(log.Fields{
		"height": height,
	})

	nextHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		logger.Error("failed to get block header")
		return typeserr.ErrGetBlockHeader
	}

	// 3. get nextHeight block header from source chain
	var header tibctypes.Header
	switch channel.source.ChainType() {
	case constant.Tendermint:
		req := &repostitory.GetBlockHeaderReq{
			LatestHeight:  nextHeight,
			TrustedHeight: height,
		}
		header, err = channel.source.GetBlockHeader(req)
		if err != nil {
			logger.Error("failed to get block header")
			return typeserr.ErrGetBlockHeader
		}
	}

	// 4. update client to dest chain
	if err := channel.dest.UpdateClient(header, channel.source.ChainName()); err != nil {
		logger.Error("failed to update client")
		return typeserr.ErrUpdateClient
	}

	return nil
}

func (channel *Channel) Relay() error {

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
		channel.logger.Error("failed to get packets")
		return typeserr.ErrGetPackets
	}
	logger := channel.logger.WithFields(log.Fields{
		"source_height": channel.context.Height(),
		"source_chain":  channel.source.ChainName(),
		"dest_chain":    channel.dest.ChainName(),
	})

	// 2.  Process biz packets
	var recvPackets types.Msgs
	for _, pack := range packets.BizPackets {

		// 2.1 get commitments packets from source chain
		commitmentsPacketResp, err := channel.source.GetCommitmentsPacket(channel.dest.ChainName(), pack.Sequence)
		if err != nil {
			logger.Error("failed to get commitment packet")
			return typeserr.ErrGetCommitmentPacket
		}

		// 2.2 get ack packet from dest chain
		ackPacketResp, err := channel.dest.GetAckPacket(channel.source.ChainName(), pack.Sequence)
		if err != nil {
			logger.Error("failed to get ack packet")
			return typeserr.ErrGetAckPacket
		}

		// 	// if ack exist, skip
		if ackPacketResp != nil {
			logger.Info("ack exist, skip cur height ")
			continue
		}

		// 2.3 get receipt packet from dest chain
		receiptPacketResp, err := channel.dest.GetReceiptPacket(channel.source.ChainName(), pack.Sequence)
		if err != nil {
			logger.Error("failed to get receipt packet")
			return typeserr.ErrGetReceiptPacket
		}
		// if receipt exist, skip
		if receiptPacketResp.Received {
			logger.Info("receipt exist, skip cur height ")
			continue
		}

		recvPacket := &packet.MsgRecvPacket{
			Packet:          pack,
			ProofCommitment: commitmentsPacketResp.Proof,
			ProofHeight: client.Height{
				RevisionNumber: commitmentsPacketResp.ProofHeight.RevisionNumber,
				RevisionHeight: commitmentsPacketResp.ProofHeight.RevisionHeight,
			},
		}
		recvPackets = append(recvPackets, recvPacket)
	}

	clientState, err := channel.dest.GetLightClientState(channel.source.ChainName())
	if err != nil {
		channel.logger.Error("failed to get light client state")
		return typeserr.ErrGetLightClientState
	}

	//3. Process ack packets
	for _, pack := range packets.AckPackets {
		// 3.1 get ack packet from dest chain
		ackPacketResp, err := channel.dest.GetAckPacket(channel.source.ChainName(), pack.Sequence)
		if err != nil {
			logger.Error("failed to get ack packet")
			return typeserr.ErrGetAckPacket
		}

		// 	// if ack exist, skip
		if ackPacketResp != nil {
			logger.Info("ack exist, skip cur height ")
			continue
		}

		// query proof
		proofHeight := channel.context.Height() + 1
		proof, err := channel.source.GetProof(channel.dest.ChainName(), pack.Sequence, proofHeight)
		if err != nil {
			logger.Error("failed to get proof")
			return typeserr.ErrGetProof
		}
		recvPacket := &packet.MsgRecvPacket{
			Packet:          pack,
			ProofCommitment: proof,
			ProofHeight: client.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: proofHeight,
			},
		}
		recvPackets = append(recvPackets, recvPacket)
	}

	for _, pack := range packets.CleanPackets {

		// query proof
		proofHeight := channel.context.Height() + 1
		proof, err := channel.source.GetProof(channel.dest.ChainName(), pack.Sequence, proofHeight)
		if err != nil {
			logger.Error("failed to get proof")
			return typeserr.ErrGetProof
		}

		recvPacket := &packet.MsgRecvPacket{
			Packet:          pack,
			ProofCommitment: proof,
			ProofHeight: client.Height{
				RevisionNumber: clientState.GetLatestHeight().GetRevisionNumber(),
				RevisionHeight: proofHeight,
			},
		}
		recvPackets = append(recvPackets, recvPacket)
	}

	// boastCommit tx
	resultTx, err := channel.dest.RecvPackets(recvPackets)
	if err != nil {
		logger.Error("failed to recv packet")
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
