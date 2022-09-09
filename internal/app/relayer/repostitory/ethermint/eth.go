package ethermint

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/irisnet/core-sdk-go/common/codec"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	"github.com/bianjieai/tibc-relayer-go/tools"

	tibcclient "github.com/bianjieai/tibc-sdk-go/client"

	"github.com/irisnet/core-sdk-go/types"

	commitmenttypes "github.com/bianjieai/tibc-sdk-go/commitment"
	tibctendermint "github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/ethermint/contracts"
	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	"github.com/bianjieai/tibc-sdk-go/packet"
	geth "github.com/ethereum/go-ethereum"
	gethcmn "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const CtxTimeout = 10 * time.Second
const TryGetGasPriceTimeInterval = 10 * time.Second

const evmStoreKey = "evm"

const (
	prefixCode = iota + 1
	prefixStorage
)

var (
	KeyPrefixStorage = []byte{prefixStorage}
)

var (
	Uint64, _  = abi.NewType("uint64", "", nil)
	Bytes32, _ = abi.NewType("bytes32", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	String, _  = abi.NewType("string", "", nil)
)

func (eth *Ethermint) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {

	pkConstr := tools.NewProofKeyConstructor(sourChainName, destChainName, sequence)
	var key []byte
	switch typ {
	case repotypes.CommitmentPoof:
		key = pkConstr.GetPacketCommitmentProofKey(eth.slot)
	case repotypes.AckProof:
		key = pkConstr.GetAckProofKey(eth.slot)
	case repotypes.CleanProof:
		key = pkConstr.GetCleanPacketCommitmentProofKey(eth.slot)
	default:
		return nil, errors.ErrGetProof
	}

	address := gethcmn.HexToAddress(eth.contractCfgGroup.Packet.Addr)

	_, proofBz, _, err := eth.getProof(int64(height), evmStoreKey, eth.stateKey(address, key))
	if err != nil {
		return nil, err
	}

	return proofBz, nil
}

func (eth *Ethermint) GetLightClientState(chainName string) (tibctypes.ClientState, error) {
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

func (eth *Ethermint) GetLightClientConsensusState(string, uint64) (tibctypes.ConsensusState, error) {
	return nil, nil
}

func (eth *Ethermint) GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error {

	hashBytes, err := eth.contracts.Packet.Commitments(nil,
		packet.PacketCommitmentKey(sourChainName, destChainName, sequence))
	if err != nil {
		return err
	}
	expectByte := make([]byte, 32)
	if bytes.Equal(expectByte, hashBytes[:]) {
		return fmt.Errorf("commitment does not exist")
	}
	return nil
}

func (eth *Ethermint) GetReceiptPacket(sourChainName, destChainName string, sequence uint64) (bool, error) {
	result, err := eth.contracts.Packet.Receipts(nil,
		packet.PacketReceiptKey(sourChainName, destChainName, sequence))
	if err != nil {
		return result, err
	}
	return result, nil
}

func (eth *Ethermint) RecvPackets(msgs types.Msgs) (*repotypes.ResultTx, types.Error) {
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

			err := eth.setPacketOpts()
			if err != nil {
				return nil, types.Wrap(err)
			}
			result, err := eth.contracts.Packet.RecvPacket(
				eth.bindOpts.packetTransactOpts,
				tmpPack,
				msg.ProofCommitment,
				height)
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

			err := eth.setPacketOpts()
			if err != nil {
				return nil, types.Wrap(err)
			}

			result, err := eth.contracts.Packet.AcknowledgePacket(
				eth.bindOpts.packetTransactOpts,
				tmpPack, msg.Acknowledgement, msg.ProofAcked,
				height)
			if err != nil {
				return nil, types.Wrap(err)
			}
			resultTx.GasUsed += int64(result.Gas())
			resultTx.Hash = resultTx.Hash + "," + result.Hash().String()
		case "recv_clean_packet":
			msg := d.(*packet.MsgRecvCleanPacket)
			cleanPack := contracts.PacketTypesCleanPacket{
				Sequence:    msg.CleanPacket.Sequence,
				DestChain:   msg.CleanPacket.DestinationChain,
				SourceChain: msg.CleanPacket.SourceChain,
				RelayChain:  msg.CleanPacket.RelayChain,
			}

			err := eth.setPacketOpts()
			if err != nil {
				return nil, types.Wrap(err)
			}

			result, err := eth.contracts.Packet.CleanPacket(
				eth.bindOpts.packetTransactOpts,
				cleanPack,
			)
			if err != nil {
				return nil, types.Wrap(err)
			}
			resultTx.GasUsed += int64(result.Gas())
			resultTx.Hash = resultTx.Hash + "," + result.Hash().String()
		}

	}

	resultTx.Hash = strings.Trim(resultTx.Hash, ",")

	return resultTx, nil
}

func (eth *Ethermint) UpdateClient(header tibctypes.Header, chainName string) (string, error) {
	h := header.(*tibctendermint.Header)
	args := abi.Arguments{
		abi.Argument{Type: Uint64},
		abi.Argument{Type: Uint64},
		abi.Argument{Type: Uint64},
		abi.Argument{Type: Bytes32},
		abi.Argument{Type: Bytes32},
	}
	timestamp := uint64(h.GetTime().Unix())
	revisionNumber := h.GetHeight().GetRevisionNumber()
	revisionHeight := h.GetHeight().GetRevisionHeight()

	var appHash [32]byte
	copy(appHash[:], h.GetHeader().AppHash[:32])

	var nextValidatorsHash [32]byte
	copy(nextValidatorsHash[:], h.GetHeader().NextValidatorsHash[:32])

	headerBytes, err := args.Pack(
		&revisionNumber,
		&revisionHeight,
		&timestamp,
		appHash,
		nextValidatorsHash,
	)

	err = eth.setClientOpts()
	if err != nil {
		return "", err
	}

	result, err := eth.contracts.Client.UpdateClient(eth.bindOpts.client, chainName, headerBytes)
	if err != nil {
		return "", err
	}

	return result.Hash().String(), nil
}

func (eth *Ethermint) GetPackets(height uint64, destChainType string) (*repotypes.Packets, error) {

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

// get packets from block
func (eth *Ethermint) getPackets(height uint64) ([]packet.Packet, error) {
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
func (eth *Ethermint) getAckPackets(height uint64) ([]repotypes.AckPacket, error) {
	address := gethcmn.HexToAddress(eth.contractCfgGroup.AckPacket.Addr)
	topic := eth.contractCfgGroup.AckPacket.Topic
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

// get clean packets from block
func (eth *Ethermint) getCleanPacket(height uint64) ([]packet.CleanPacket, error) {
	address := gethcmn.HexToAddress(eth.contractCfgGroup.AckPacket.Addr)
	topic := eth.contractCfgGroup.CleanPacket.Topic
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

func (eth *Ethermint) getLogs(address gethcmn.Address, topic string, fromBlock, toBlock uint64) ([]gethtypes.Log, error) {
	filter := geth.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []gethcmn.Address{address},
		Topics:    [][]gethcmn.Hash{{gethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return eth.ethClient.FilterLogs(context.Background(), filter)
}

func (eth *Ethermint) getGasPrice() (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return eth.ethClient.SuggestGasPrice(ctx)

}

func (eth *Ethermint) setPacketOpts() error {
	var curGasPrice *big.Int
	for {
		gasPrice, err := eth.getGasPrice()
		if err != nil {
			return err
		}
		cmpRes := eth.maxGasPrice.Cmp(gasPrice)
		if cmpRes == -1 {
			time.Sleep(TryGetGasPriceTimeInterval)
			continue
		} else {
			gasPriceUint := gasPrice.Int64()
			gasPriceUint += int64(float64(gasPriceUint) * eth.tipCoefficient)
			curGasPrice = new(big.Int).SetInt64(gasPriceUint)
			break
		}
	}

	eth.bindOpts.packetTransactOpts.GasPrice = curGasPrice
	return nil
}

func (eth *Ethermint) setClientOpts() error {
	var curGasPrice *big.Int
	for {
		gasPrice, err := eth.getGasPrice()
		if err != nil {
			return err
		}
		cmpRes := eth.maxGasPrice.Cmp(gasPrice)
		if cmpRes == -1 {
			continue
		} else {
			gasPriceUint := gasPrice.Int64()
			gasPriceUint += int64(float64(gasPriceUint) * eth.tipCoefficient)
			curGasPrice = new(big.Int).SetInt64(gasPriceUint)
			break
		}
	}

	eth.bindOpts.client.GasPrice = curGasPrice
	return nil
}

func (eth *Ethermint) getProof(height int64, storeKey string, key []byte) ([]byte, []byte, uint64, error) {
	// ABCI queries at heights 1, 2 or less than or equal to 0 are not supported.
	// Base app does not support queries for height less than or equal to 1.
	// Therefore, a query at height 2 would be equivalent to a query at height 3.
	// A height of 0 will query with the lastest state.
	if height != 0 && height <= 2 {
		return nil, nil, 0, fmt.Errorf("proof queries at height <= 2 are not supported")
	}

	// Use the IAVL height if a valid tendermint height is passed in.
	// A height of 0 will query with the latest state.
	if height != 0 {
		height--
	}

	res, err := eth.terndermintCli.QueryStore(key, storeKey, height, true)
	if err != nil {
		return nil, nil, 0, err
	}

	merkleProof, err := commitmenttypes.ConvertProofs(res.ProofOps)
	if err != nil {
		return nil, nil, 0, err
	}

	cdc := codec.NewProtoCodec(eth.terndermintCli.TIBC.InterfaceRegistry)

	proofBz, err := cdc.MarshalBinaryBare(&merkleProof)
	if err != nil {
		return nil, nil, 0, err
	}
	return res.Value, proofBz, uint64(res.Height) + 1, nil
}

// StateKey defines the full key under which an account state is stored.
func (eth *Ethermint) stateKey(address common.Address, key []byte) []byte {
	return append(addressStoragePrefix(address), key...)
}

// AddressStoragePrefix returns a prefix to iterate over a given account storage.
func addressStoragePrefix(address common.Address) []byte {
	return append(KeyPrefixStorage, address.Bytes()...)
}
