package repostitory

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"

	"github.com/tendermint/tendermint/light/provider"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/irisnet/core-sdk-go/bank"
	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/gov"
	"github.com/irisnet/core-sdk-go/staking"
	"github.com/irisnet/irismod-sdk-go/nft"

	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	tibcnfttypes "github.com/bianjieai/tibc-sdk-go/nft_transfer"
	"github.com/bianjieai/tibc-sdk-go/packet"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/common/codec"
	cdctypes "github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	"github.com/irisnet/core-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"
	"github.com/tendermint/tendermint/libs/log"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

var _ IChain = new(Tendermint)

var (
	maxRetryAttempts    = 5
	regexpTooHigh       = regexp.MustCompile(`height \d+ must be less than or equal to`)
	regexpMissingHeight = regexp.MustCompile(`height \d+ is not available`)
	regexpTimedOut      = regexp.MustCompile(`Timeout exceeded`)
)

type Tendermint struct {
	logger log.Logger

	terndermintCli tendermintClient
	baseTx         types.BaseTx
	address        string

	chainName             string
	chainType             string
	revisionNumber        uint64
	updateClientFrequency uint64

	allowMapSender     map[string][]string
	cleanPacketEnabled bool
}

func NewTendermintClient(
	chainType string,
	chainName string,
	updateClientFrequency uint64,
	allowMapSender map[string][]string,
	cleanPacketEnabled bool,
	config *TerndermintConfig) (*Tendermint, error) {
	cfg, err := coretypes.NewClientConfig(config.RPCAddr, config.GrpcAddr, config.ChainID,
		config.Options...)
	if err != nil {
		return nil, err
	}
	tc := newTendermintClient(cfg, chainName)

	// import key to core-sdk
	address, err := tc.BaseClient.Import(config.Name, config.Password, config.PrivKeyArmor)
	if err != nil {
		return nil, err
	}
	fmt.Println(address)

	return &Tendermint{
		chainType:             chainType,
		chainName:             chainName,
		terndermintCli:        tc,
		updateClientFrequency: updateClientFrequency,
		logger:                tc.BaseClient.Logger(),
		baseTx:                config.BaseTx,
		address:               address,
		allowMapSender:        allowMapSender,
		cleanPacketEnabled:    cleanPacketEnabled,
	}, err
}

func (c *Tendermint) GetPackets(height uint64, destChainType string) (*repotypes.Packets, error) {
	var bizPackets []packet.Packet
	var ackPackets []repotypes.AckPacket
	var cleanPackets []packet.CleanPacket

	curHeight := int64(height)
	block, err := c.terndermintCli.Block(context.Background(), &curHeight)
	if err != nil {
		return nil, err
	}

	packets := repotypes.NewPackets()

	for _, tx := range block.Block.Txs {
		hash := hex.EncodeToString(tx.Hash())
		resultTx, err := c.terndermintCli.BaseClient.QueryTx(hash)
		if err != nil {
			//todo
			// 需要修改sdk
			continue
		}
		if c.isExistPacket(repotypes.EventTypeSendPacket, resultTx) {
			tmpPacket, err := c.getPacket(resultTx, destChainType)
			if err != nil {
				return nil, err
			}
			bizPackets = append(bizPackets, tmpPacket...)
		}

		if c.isExistPacket(repotypes.EventTypeWriteAck, resultTx) {
			// get ack packet
			tmpAckPacks, acks, err := c.getAckPackets(resultTx, destChainType)
			if err != nil {
				return nil, err
			}
			for i := 0; i < len(tmpAckPacks); i++ {
				tmpAckPacket := repotypes.AckPacket{
					Packet:          tmpAckPacks[i],
					Acknowledgement: acks[i],
				}
				ackPackets = append(ackPackets, tmpAckPacket)
			}

		}

		if c.cleanPacketEnabled {
			if c.isExistPacket(repotypes.EventTypeSendCleanPacket, resultTx) {
				tmpCleanPacket, err := c.getCleanPacket(resultTx, destChainType)
				if err != nil {
					return nil, err
				}
				cleanPackets = append(cleanPackets, tmpCleanPacket...)
			}
		}

	}

	packets.BizPackets = bizPackets
	packets.AckPackets = ackPackets
	packets.CleanPackets = cleanPackets
	return packets, nil
}

func (c *Tendermint) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {
	var key []byte
	switch typ {
	case repotypes.CommitmentPoof:
		key = packet.PacketCommitmentKey(sourChainName, destChainName, sequence)
	case repotypes.AckProof:
		key = packet.PacketAcknowledgementKey(sourChainName, destChainName, sequence)
	case repotypes.CleanProof:
		key = packet.CleanPacketCommitmentKey(sourChainName, destChainName)
	default:
		return nil, errors.ErrGetProof
	}

	_, proofBz, _, err := c.terndermintCli.TIBC.QueryTendermintProof(int64(height), key)
	if err != nil {
		return nil, err
	}
	return proofBz, nil
}

func (c *Tendermint) RecvPackets(msgs types.Msgs) (*repotypes.ResultTx, types.Error) {
	for _, d := range msgs {
		switch d.Type() {
		case "recv_packet":
			msg := d.(*packet.MsgRecvPacket)
			msg.Signer = c.address
		case "acknowledge_packet":
			msg := d.(*packet.MsgAcknowledgement)
			msg.Signer = c.address
		case "recv_clean_packet":
			msg := d.(*packet.MsgRecvCleanPacket)
			msg.Signer = c.address
		}
	}

	resultTx, err := c.terndermintCli.TIBC.RecvPackets(msgs, c.baseTx)
	if err != nil {
		return nil, types.Wrap(err)
	}
	return &repotypes.ResultTx{
		GasWanted: resultTx.GasWanted,
		GasUsed:   resultTx.GasUsed,
		Hash:      resultTx.Hash,
		Height:    resultTx.Height,
	}, nil
}

func (c *Tendermint) GetBlockTimestamp(height uint64) (uint64, error) {
	block, err := c.terndermintCli.QueryBlock(int64(height))
	if err != nil {
		return 0, err
	}
	return uint64(block.Block.Time.Unix()), nil
}

func (c *Tendermint) GetBlockHeader(req *repotypes.GetBlockHeaderReq) (tibctypes.Header, error) {
	block, err := c.terndermintCli.QueryBlock(int64(req.LatestHeight))
	if err != nil {
		return nil, err
	}
	rescommit, err := c.terndermintCli.Commit(context.Background(), &block.BlockResult.Height)
	if err != nil {
		return nil, err
	}
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: block.Block.Header.ToProto(),
		Commit: commit.ToProto(),
	}

	validatorSet, err := c.getValidator(int64(req.LatestHeight))
	if err != nil {
		return nil, err

	}
	trustedValidators, err := c.getValidator(int64(req.TrustedHeight))
	if err != nil {
		return nil, err
	}
	// The trusted fields may be nil. They may be filled before relaying messages to a client.
	// The relayer is responsible for querying client and injecting appropriate trusted fields.
	return &tendermint.Header{
		SignedHeader: signedHeader,
		ValidatorSet: validatorSet,
		TrustedHeight: tibcclient.Height{
			RevisionNumber: req.RevisionNumber,
			RevisionHeight: req.TrustedHeight,
		},
		TrustedValidators: trustedValidators,
	}, nil
}

func (c *Tendermint) GetLightClientState(chainName string) (tibctypes.ClientState, error) {
	return c.terndermintCli.TIBC.GetClientState(chainName)

}

func (c *Tendermint) GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {
	return c.terndermintCli.TIBC.GetConsensusState(chainName, height)

}

func (c *Tendermint) GetLatestHeight() (uint64, error) {
	block, err := c.terndermintCli.Block(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	var height = block.Block.Height
	return uint64(height), err
}

func (c *Tendermint) GetResult(hash string) (uint64, error) {
	res, err := c.terndermintCli.QueryTx(hash)
	if err != nil {
		return 0, err
	}
	code := uint64(res.Result.Code)
	return code, nil
}

func (c *Tendermint) GetLightClientDelayHeight(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayBlock(), nil
}

func (c *Tendermint) GetLightClientDelayTime(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayTime(), nil

}

func (c *Tendermint) UpdateClient(header tibctypes.Header, chainName string) (string, error) {
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    header,
	}
	resTx, err := c.terndermintCli.TIBC.UpdateClient(request, c.baseTx)
	if err != nil {
		return "", err
	}
	return resTx.Hash, nil
}

func (c *Tendermint) GetCommitmentsPacket(sourceChainName, destChainName string, sequence uint64) error {
	_, err := c.terndermintCli.TIBC.PacketCommitment(destChainName, sourceChainName, sequence)
	if err != nil {
		return err
	}
	return nil
}

func (c *Tendermint) GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (bool, error) {
	result, err := c.terndermintCli.TIBC.PacketReceipt(destChianName, sourChainName, sequence)
	if err != nil {
		return false, err
	}
	return result.Received, nil
}

func (c *Tendermint) ChainName() string {

	return c.chainName
}

func (c *Tendermint) ChainType() string {
	return c.chainType
}

func (c *Tendermint) UpdateClientFrequency() uint64 {
	return c.updateClientFrequency
}

func (c *Tendermint) getValidator(height int64) (*tenderminttypes.ValidatorSet, error) {
	const maxPages = 100

	var (
		perPage = 100
		vals    = []*tmtypes.Validator{}
		page    = 1
		total   = -1
	)
	ctx := context.Background()

OUTER_LOOP:
	for len(vals) != total && page <= maxPages {
		for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
			res, err := c.terndermintCli.TIBC.Validators(ctx, &height, &page, &perPage)
			switch {
			case err == nil:
				// Validate response.
				if len(res.Validators) == 0 {
					return nil, provider.ErrBadLightBlock{
						Reason: fmt.Errorf("validator set is empty (height: %d, page: %d, per_page: %d)",
							height, page, perPage),
					}
				}
				if res.Total <= 0 {
					return nil, provider.ErrBadLightBlock{
						Reason: fmt.Errorf("total number of vals is <= 0: %d (height: %d, page: %d, per_page: %d)",
							res.Total, height, page, perPage),
					}
				}

				total = res.Total
				vals = append(vals, res.Validators...)
				page++
				continue OUTER_LOOP

			case regexpTooHigh.MatchString(err.Error()):
				return nil, fmt.Errorf("height requested is too high")

			case regexpMissingHeight.MatchString(err.Error()):
				return nil, provider.ErrLightBlockNotFound

			// if we have exceeded retry attempts then return no response error
			case attempt == maxRetryAttempts:
				return nil, provider.ErrNoResponse

			case regexpTimedOut.MatchString(err.Error()):
				// we wait and try again with exponential backoff
				time.Sleep(backoffTimeout(uint16(attempt)))
				continue

			// context canceled or connection refused we return the error
			default:
				return nil, err
			}

		}
	}
	validatorSet, err := tmtypes.NewValidatorSet(vals).ToProto()
	if err != nil {
		return nil, err
	}

	return validatorSet, nil
}

// exponential backoff (with jitter)
// 0.5s -> 2s -> 4.5s -> 8s -> 12.5 with 1s variation
func backoffTimeout(attempt uint16) time.Duration {
	// nolint:gosec // G404: Use of weak random number generator
	return time.Duration(500*attempt*attempt)*time.Millisecond + time.Duration(rand.Intn(1000))*time.Millisecond
}

func (c *Tendermint) getPacket(tx types.ResultQueryTx, destChainType string) ([]packet.Packet, error) {
	sequences := tx.Result.Events.GetValues(repotypes.EventTypeSendPacket, "packet_sequence")
	srcChains := tx.Result.Events.GetValues(repotypes.EventTypeSendPacket, "packet_src_chain")
	dstPorts := tx.Result.Events.GetValues(repotypes.EventTypeSendPacket, "packet_dst_port")
	ports := tx.Result.Events.GetValues(repotypes.EventTypeSendPacket, "packet_port")
	rlyChains := tx.Result.Events.GetValues(repotypes.EventTypeSendPacket, "packet_relay_channel")
	datas := tx.Result.Events.GetValues(repotypes.EventTypeSendPacket, "packet_data")

	var packets []packet.Packet
	for i := 0; i < len(sequences); i++ {
		sequenceStr := sequences[i]
		sequence, err := strconv.Atoi(sequenceStr)
		if err != nil {
			return nil, err
		}

		tmpPack := packet.Packet{
			Sequence:         uint64(sequence),
			SourceChain:      srcChains[i],
			DestinationChain: dstPorts[i],
			Port:             ports[i],
			RelayChain:       rlyChains[i],
			Data:             []byte(datas[i]),
		}

		nonFungibleTokenPacketData := &tibcnfttypes.NonFungibleTokenPacketData{}
		if err := nonFungibleTokenPacketData.Unmarshal(tmpPack.Data); err != nil {
			continue
		}
		//msgNftTransfer.DestContract

		//allowList set
		//msg.sender not in allowList, skip
		senders, ok := c.allowMapSender[nonFungibleTokenPacketData.DestContract]
		if ok && !c.isExitsFromStringList(senders, nonFungibleTokenPacketData.Sender) {
			continue
		}

		// 1. iris to iris?
		// 2. iris to eth?

		if destChainType == constant.ETH && !ok {
			// if nonFungibleTokenPacketData.DestContract not in allowList, skip
			continue
		}

		packets = append(packets, tmpPack)
	}

	return packets, nil
}

func (c *Tendermint) getAckPackets(tx types.ResultQueryTx, destChainType string) ([]packet.Packet, [][]byte, error) {

	sequences := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_sequence")
	srcChains := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_src_chain")
	dstPorts := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_dst_port")
	ports := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_port")
	rlyChains := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_relay_channel")
	datas := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_data")
	acks := tx.Result.Events.GetValues(repotypes.EventTypeWriteAck, "packet_ack")
	var ackByteList [][]byte
	var packets []packet.Packet
	for i := 0; i < len(sequences); i++ {
		sequenceStr := sequences[i]
		sequence, err := strconv.Atoi(sequenceStr)
		if err != nil {
			return nil, nil, err
		}
		tmpAckPack := packet.Packet{
			Sequence:         uint64(sequence),
			SourceChain:      srcChains[i],
			DestinationChain: dstPorts[i],
			Port:             ports[i],
			RelayChain:       rlyChains[i],
			Data:             []byte(datas[i]),
		}
		ackByteList = append(ackByteList, []byte(acks[i]))
		packets = append(packets, tmpAckPack)
	}

	return packets, ackByteList, nil
}

func (c *Tendermint) getCleanPacket(tx types.ResultQueryTx, destChainType string) ([]packet.CleanPacket, error) {
	sequences := tx.Result.Events.GetValues(repotypes.EventTypeSendCleanPacket, "packet_sequence")
	sourceChains := tx.Result.Events.GetValues(repotypes.EventTypeSendCleanPacket, "packet_src_chain")
	dstPorts := tx.Result.Events.GetValues(repotypes.EventTypeSendCleanPacket, "packet_dst_port")
	rlyChains := tx.Result.Events.GetValues(repotypes.EventTypeSendCleanPacket, "packet_relay_channel")
	var packets []packet.CleanPacket
	for i := 0; i < len(sequences); i++ {
		sequenceStr := sequences[i]
		sequence, err := strconv.Atoi(sequenceStr)
		if err != nil {
			return nil, err
		}
		tmpCleanPack := packet.CleanPacket{
			Sequence:         uint64(sequence),
			SourceChain:      sourceChains[i],
			DestinationChain: dstPorts[i],
			RelayChain:       rlyChains[i],
		}
		packets = append(packets, tmpCleanPack)
	}

	return packets, nil

}

func (c *Tendermint) isExistPacket(typ string, tx types.ResultQueryTx) bool {
	_, err := tx.Result.Events.GetValue(typ, "packet_sequence")
	if err != nil {
		return false
	}
	return true
}

func (c *Tendermint) isExitsFromStringList(sources []string, target string) bool {
	for _, source := range sources {
		if source == target {
			return true
		}
	}
	return false
}

//======================================

type tendermintClient struct {
	encodingConfig types.EncodingConfig
	coretypes.BaseClient
	Bank      bank.Client
	Staking   staking.Client
	Gov       gov.Client
	NFT       nft.Client
	TIBC      tibc.Client
	ChainName string
}

func newTendermintClient(cfg types.ClientConfig, chainName string) tendermintClient {
	encodingConfig := makeEncodingConfig()
	// create a instance of baseClient
	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	stakingClient := staking.NewClient(baseClient, encodingConfig.Marshaler)
	govClient := gov.NewClient(baseClient, encodingConfig.Marshaler)
	tibcClient := tibc.NewClient(baseClient, encodingConfig)
	nftClient := nft.NewClient(baseClient, encodingConfig.Marshaler)

	tc := &tendermintClient{
		encodingConfig: encodingConfig,
		BaseClient:     baseClient,
		Bank:           bankClient,
		Staking:        stakingClient,
		Gov:            govClient,
		NFT:            nftClient,
		TIBC:           tibcClient,
		ChainName:      chainName,
	}

	tc.RegisterModule(
		bankClient,
		stakingClient,
		govClient,
	)
	return *tc
}

func (tc tendermintClient) Manager() types.BaseClient {
	return tc.BaseClient
}

func (tc tendermintClient) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(tc.encodingConfig.InterfaceRegistry)
	}
}

//client init
func makeEncodingConfig() types.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	registerLegacyAminoCodec(encodingConfig.Amino)
	registerInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func registerLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func registerInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}

type TerndermintConfig struct {
	Options      []coretypes.Option
	BaseTx       types.BaseTx
	PrivKeyArmor string
	Name         string
	Password     string

	RPCAddr  string
	GrpcAddr string
	ChainID  string
}

func NewTerndermintConfig() *TerndermintConfig {
	return &TerndermintConfig{}
}
