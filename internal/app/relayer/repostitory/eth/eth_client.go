package eth

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/eth/contracts"
	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	"github.com/bianjieai/tibc-relayer-go/tools"

	geth "github.com/ethereum/go-ethereum"
	gethcmn "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	tibceth "github.com/bianjieai/tibc-sdk-go/eth"
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctendermint "github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/types"
)

var _ repostitory.IChain = new(Eth)

const CtxTimeout = 10 * time.Second

type Eth struct {
	uri                   string
	chainName             string
	chainType             string
	updateClientFrequency uint64

	contractCfgGroup *ContractCfgGroup
	contracts        *filter
	bindOpts         *bindOpts

	ethClient *gethethclient.Client
	gethCli   *gethclient.Client

	amino *codec.LegacyAmino
}

func NewEth(chainType, chainName string, updateClientFrequency uint64, uri string, cfgGroup *ContractCfgGroup) (repostitory.IChain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, uri)
	if err != nil {
		return nil, err
	}
	ethClient := gethethclient.NewClient(rpcClient)
	gethCli := gethclient.New(rpcClient)

	filter, err := newFilter(ethClient, cfgGroup)
	if err != nil {
		return nil, err
	}

	tmpBindOpts, err := newBindOpts(cfgGroup.Packet.OptPrivKey, cfgGroup.Client.OptPrivKey)

	if err != nil {
		return nil, err
	}

	codecAmino := codec.NewLegacyAmino()

	return &Eth{
		chainType:             chainType,
		chainName:             chainName,
		updateClientFrequency: updateClientFrequency,
		contractCfgGroup:      cfgGroup,
		ethClient:             ethClient,
		gethCli:               gethCli,
		contracts:             filter,
		amino:                 codecAmino,
		bindOpts:              tmpBindOpts,
	}, nil
}

func (eth *Eth) RecvPackets(msgs types.Msgs) (*repotypes.ResultTx, types.Error) {
	resultTx := &repotypes.ResultTx{}

	for _, d := range msgs {

		switch d.Type() {
		case "recv_packet":
			msg := d.(*packet.MsgRecvPacket)

			tmpPack := contracts.PacketTypesPacket{
				Sequence:    msg.Packet.Sequence,
				Port:        msg.Packet.Port,
				DestChain:   msg.Packet.DestinationChain,
				SourceChain: msg.Packet.SourceChain,
				RelayChain:  msg.Packet.RelayChain,
				Data:        msg.Packet.Data,
			}
			height := contracts.HeightData{
				RevisionNumber: msg.ProofHeight.RevisionNumber,
				RevisionHeight: msg.ProofHeight.RevisionHeight,
			}
			result, err := eth.contracts.Packet.RecvPacket(eth.bindOpts.packet, tmpPack, msg.ProofCommitment, height)
			if err != nil {
				return nil, types.Wrap(err)
			}
			resultTx.GasUsed += int64(result.Gas())
			resultTx.Hash = resultTx.Hash + "," + result.Hash().String()
		case "acknowledge_packet":
			msg := d.(*packet.MsgAcknowledgement)
			tmpPack := contracts.PacketTypesPacket{
				Sequence:    msg.Packet.Sequence,
				Port:        msg.Packet.Port,
				DestChain:   msg.Packet.DestinationChain,
				SourceChain: msg.Packet.SourceChain,
				RelayChain:  msg.Packet.RelayChain,
				Data:        msg.Packet.Data,
			}
			height := contracts.HeightData{
				RevisionNumber: msg.ProofHeight.RevisionNumber,
				RevisionHeight: msg.ProofHeight.RevisionHeight,
			}
			result, err := eth.contracts.Packet.RecvPacket(nil, tmpPack, msg.ProofAcked, height)
			if err != nil {
				return nil, types.Wrap(err)
			}
			resultTx.GasUsed += int64(result.Gas())
			resultTx.Hash = resultTx.Hash + "," + result.Hash().String()
		}

	}

	return resultTx, nil
}

func (eth *Eth) UpdateClient(header tibctypes.Header, chainName string) (string, error) {
	h := header.(*tibctendermint.Header)
	headerBytes, err := h.Marshal()
	if err != nil {
		return "", err
	}
	fmt.Println(hex.EncodeToString(headerBytes))
	result, err := eth.contracts.Client.UpdateClient(eth.bindOpts.client, chainName, headerBytes)
	if err != nil {
		return "", err
	}

	return result.Hash().String(), nil
}

func (eth *Eth) GetPackets(height uint64) (*repotypes.Packets, error) {

	bizPackets, err := eth.getPackets(height)
	if err != nil {
		return nil, err
	}
	ackPackets, err := eth.getAckPackets(height)
	if err != nil {
		return nil, err
	}
	cleanPackets, err := eth.getCleanPacket(height)
	if err != nil {
		return nil, err
	}

	packets := &repotypes.Packets{
		BizPackets:   bizPackets,
		AckPackets:   ackPackets,
		CleanPackets: cleanPackets,
	}

	return packets, nil
}

func (eth *Eth) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {
	pkConstr := tools.NewProofKeyConstructor(sourChainName, destChainName, sequence)
	var key []byte
	switch typ {
	case repotypes.CommitmentPoof:
		key = pkConstr.GetPacketCommitmentProofKey()
	case repotypes.AckProof:
		key = pkConstr.GetAckProofKey()
	default:
		return nil, errors.ErrGetProof
	}
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	address := gethcmn.HexToAddress(eth.contractCfgGroup.Packet.Addr)
	result, err := eth.gethCli.GetProof(ctx, address, []string{string(key)}, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}
	proofBz, err := eth.amino.MarshalBinaryBare(result)
	if err != nil {
		return nil, err
	}
	return proofBz, nil
}

func (eth *Eth) GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error {

	hashBytes, err := eth.contracts.Packet.Commitments(nil, packet.PacketCommitmentKey(sourChainName, destChainName, sequence))
	if err != nil {
		return err
	}
	expectByte := make([]byte, 32)
	if bytes.Equal(expectByte, hashBytes[:]) {
		return fmt.Errorf("commitment does not exist")
	}
	return nil
}

func (eth *Eth) GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (bool, error) {
	result, err := eth.contracts.Packet.Receipts(nil, []byte(""))
	if err != nil {
		return false, err
	}
	return result, nil
}

func (eth *Eth) GetBlockTimestamp(height uint64) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := eth.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return 0, err
	}
	return blockRes.Time(), nil
}

func (eth *Eth) GetBlockHeader(req *repotypes.GetBlockHeaderReq) (tibctypes.Header, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := eth.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(req.LatestHeight))
	if err != nil {
		return nil, err
	}

	return &tibceth.Header{
		ParentHash:  blockRes.ParentHash().Bytes(),
		UncleHash:   blockRes.UncleHash().Bytes(),
		Coinbase:    blockRes.Coinbase().Bytes(),
		Root:        blockRes.Root().Bytes(),
		TxHash:      blockRes.TxHash().Bytes(),
		ReceiptHash: blockRes.ReceiptHash().Bytes(),
		Bloom:       blockRes.Bloom().Bytes(),
		Difficulty:  blockRes.Difficulty().Uint64(),
		Height: tibcclient.Height{
			RevisionNumber: 0,
			RevisionHeight: req.TrustedHeight,
		},
		GasLimit:  blockRes.GasLimit(),
		GasUsed:   blockRes.GasUsed(),
		Time:      blockRes.Time(),
		Extra:     blockRes.Extra(),
		MixDigest: blockRes.MixDigest().Bytes(),
		Nonce:     blockRes.Nonce(),
		BaseFee:   blockRes.BaseFee().Uint64(),
	}, nil

}

func (eth *Eth) GetLightClientState(chainName string) (tibctypes.ClientState, error) {
	latestHeight, err := eth.contracts.Client.GetLatestHeight(nil, chainName)
	if err != nil {
		return nil, err
	}
	return &tibctendermint.ClientState{
		LatestHeight: tibcclient.Height{
			RevisionHeight: latestHeight.RevisionHeight,
			RevisionNumber: latestHeight.RevisionNumber,
		},
	}, nil
}

func (eth *Eth) GetLightClientConsensusState(string, uint64) (tibctypes.ConsensusState, error) {
	return nil, nil
}

func (eth *Eth) GetLatestHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return eth.ethClient.BlockNumber(ctx)
}

func (eth *Eth) GetLightClientDelayHeight(chainName string) (uint64, error) {
	return 0, nil
}

func (eth *Eth) GetLightClientDelayTime(chainName string) (uint64, error) {
	return 0, nil
}

func (eth *Eth) ChainName() string {
	return eth.chainName
}
func (eth *Eth) UpdateClientFrequency() uint64 {
	return eth.updateClientFrequency
}

func (eth *Eth) ChainType() string {
	return eth.chainType
}

// get packets from block
func (eth *Eth) getPackets(height uint64) ([]packet.Packet, error) {
	address := gethcmn.HexToAddress(eth.contractCfgGroup.Packet.Addr)
	topic := eth.contractCfgGroup.Packet.Topic
	logs, err := eth.getLogs(address, topic, height, height)
	if err != nil {
		return nil, err
	}

	var bizPackets []packet.Packet
	for _, log := range logs {
		packSent, err := eth.contracts.Packet.ParsePacketSent(log)
		if err != nil {
			return nil, err
		}
		tmpPack := packet.Packet{
			Sequence:         packSent.Packet.Sequence,
			Data:             packSent.Packet.Data,
			SourceChain:      packSent.Packet.SourceChain,
			DestinationChain: packSent.Packet.DestChain,
			Port:             packSent.Packet.Port,
			RelayChain:       packSent.Packet.RelayChain,
		}
		bizPackets = append(bizPackets, tmpPack)

	}
	return bizPackets, nil
}

// get ack packets from block
func (eth *Eth) getAckPackets(height uint64) ([]repotypes.AckPacket, error) {
	address := gethcmn.HexToAddress(eth.contractCfgGroup.AckPacket.Addr)
	topic := eth.contractCfgGroup.Packet.Topic
	logs, err := eth.getLogs(address, topic, height, height)
	if err != nil {
		return nil, err
	}

	var ackPackets []repotypes.AckPacket
	for _, log := range logs {
		ackWritten, err := eth.contracts.Packet.ParseAckWritten(log)
		if err != nil {
			return nil, err
		}
		tmpAckPack := repotypes.AckPacket{}
		tmpAckPack.Packet = packet.Packet{
			Sequence:         ackWritten.Packet.Sequence,
			Data:             ackWritten.Packet.Data,
			SourceChain:      ackWritten.Packet.SourceChain,
			DestinationChain: ackWritten.Packet.DestChain,
			Port:             ackWritten.Packet.Port,
			RelayChain:       ackWritten.Packet.RelayChain,
		}
		tmpAckPack.Acknowledgement = ackWritten.Ack
		ackPackets = append(ackPackets, tmpAckPack)
	}
	return ackPackets, nil
}

func (eth *Eth) getCleanPacket(height uint64) ([]packet.CleanPacket, error) {
	address := gethcmn.HexToAddress(eth.contractCfgGroup.AckPacket.Addr)
	topic := eth.contractCfgGroup.Packet.Topic
	logs, err := eth.getLogs(address, topic, height, height)
	if err != nil {
		return nil, err
	}

	var cleanPackets []packet.CleanPacket
	for _, log := range logs {
		packSent, err := eth.contracts.Packet.ParseCleanPacketSent(log)
		if err != nil {
			return nil, err
		}
		tmpPack := packet.CleanPacket{
			Sequence:         packSent.Packet.Sequence,
			SourceChain:      packSent.Packet.SourceChain,
			DestinationChain: packSent.Packet.DestChain,
			RelayChain:       packSent.Packet.RelayChain,
		}
		cleanPackets = append(cleanPackets, tmpPack)
	}
	return cleanPackets, nil
}

func (eth *Eth) getLogs(address gethcmn.Address, topic string, fromBlock, toBlock uint64) ([]gethtypes.Log, error) {
	filter := geth.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []gethcmn.Address{address},
		Topics:    [][]gethcmn.Hash{{gethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return eth.ethClient.FilterLogs(context.Background(), filter)
}

type bindOpts struct {
	client *bind.TransactOpts
	packet *bind.TransactOpts
}

func newBindOpts(clientPrivKey, packetPrivKey string) (*bindOpts, error) {

	cliPriv, err := gethcrypto.HexToECDSA(clientPrivKey)
	if err != nil {
		return nil, err
	}
	chainID := new(big.Int).SetUint64(4)
	clientOpts, err := bind.NewKeyedTransactorWithChainID(cliPriv, chainID)
	if err != nil {
		return nil, err
	}
	var tmpPacketPrivKey string
	if len(packetPrivKey) == 0 {
		tmpPacketPrivKey = clientPrivKey
	} else {
		tmpPacketPrivKey = packetPrivKey
	}
	clientOpts.GasLimit = 20000000
	clientOpts.GasPrice = new(big.Int).SetUint64(1500000000)
	packPriv, err := gethcrypto.HexToECDSA(tmpPacketPrivKey)
	if err != nil {
		return nil, err
	}
	packOpts := bind.NewKeyedTransactor(packPriv)
	packOpts.GasLimit = 20000000
	packOpts.GasPrice = new(big.Int).SetUint64(1500000000)
	return &bindOpts{
		client: clientOpts,
		packet: packOpts,
	}, nil
}

type filter struct {
	Packet *contracts.Packet
	Client *contracts.Client
}

func newFilter(ethClient *gethethclient.Client, cfgGroup *ContractCfgGroup) (*filter, error) {
	packAddr := gethcmn.HexToAddress(cfgGroup.Packet.Addr)
	packetFilter, err := contracts.NewPacket(packAddr, ethClient)
	if err != nil {
		return nil, err
	}

	clientAddr := gethcmn.HexToAddress(cfgGroup.Client.Addr)
	clientFilter, err := contracts.NewClient(clientAddr, ethClient)
	if err != nil {
		return nil, err
	}

	return &filter{
		Packet: packetFilter,
		Client: clientFilter,
	}, nil
}

type ContractCfgGroup struct {
	Client      ContractCfg
	Packet      ContractCfg
	AckPacket   ContractCfg
	CleanPacket ContractCfg
}

type ContractCfg struct {
	Addr       string
	Topic      string
	OptPrivKey string
}

func NewContracts() *ContractCfgGroup {
	return &ContractCfgGroup{}
}
