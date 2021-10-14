package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/eth/contracts"
	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/errors"
	"github.com/bianjieai/tibc-relayer-go/tools"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	"github.com/irisnet/core-sdk-go/types"
)

var _ repostitory.IChain = new(Eth)

const CtxTimeout = 10 * time.Second

var (
	Uint64, _  = abi.NewType("uint64", "", nil)
	Bytes32, _ = abi.NewType("bytes32", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
	String, _  = abi.NewType("string", "", nil)
)

type Eth struct {
	uri                   string
	chainName             string
	chainType             string
	updateClientFrequency uint64

	contractCfgGroup *ContractCfgGroup
	contracts        *contractGroup
	bindOpts         *bindOpts

	slot int64

	ethClient  *gethethclient.Client
	gethCli    *gethclient.Client
	gethRpcCli *gethrpc.Client

	//amino codec.Marshaler
}

func NewEth(config *ChainConfig) (repostitory.IChain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, config.ChainURI)
	if err != nil {
		return nil, err
	}

	ethClient := gethethclient.NewClient(rpcClient)
	gethCli := gethclient.New(rpcClient)

	contractGroup, err := newContractGroup(ethClient, config.ContractCfgGroup)
	if err != nil {
		return nil, err
	}

	tmpBindOpts, err := newBindOpts(config.ContractBindOptsCfg)

	if err != nil {
		return nil, err
	}

	return &Eth{
		chainType:             config.ChainType,
		chainName:             config.ChainName,
		updateClientFrequency: config.UpdateClientFrequency,
		contractCfgGroup:      config.ContractCfgGroup,
		ethClient:             ethClient,
		gethCli:               gethCli,
		gethRpcCli:            rpcClient,
		contracts:             contractGroup,
		bindOpts:              tmpBindOpts,
		slot:                  config.Slot,
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

func (eth *Eth) UpdateClient(header tibctypes.Header, chainName string) (string, error) {
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
		key = pkConstr.GetPacketCommitmentProofKey(eth.slot)
	case repotypes.AckProof:
		key = pkConstr.GetAckProofKey(eth.slot)
	case repotypes.CleanProof:
		key = pkConstr.GetCleanPacketCommitmentProofKey(eth.slot)
	default:
		return nil, errors.ErrGetProof
	}
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	address := gethcmn.HexToAddress(eth.contractCfgGroup.Packet.Addr)
	//fmt.Println(hex.DecodeString(string(key)))
	result, err := eth.getProof(ctx, address, []string{hexutil.Encode(key)}, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}

	var storageProof []*tibceth.StorageResult
	for _, sp := range result.StorageProof {

		tmpStorageProof := &tibceth.StorageResult{
			Key:   sp.Key,
			Value: hexutil.EncodeBig(sp.Value),
			Proof: sp.Proof,
		}

		storageProof = append(storageProof, tmpStorageProof)
	}
	nonce := hexutil.EncodeUint64(result.Nonce)
	balance := hexutil.EncodeBig(result.Balance)
	proof := &tibceth.Proof{
		Address:      result.Address.String(),
		Balance:      balance,
		CodeHash:     result.CodeHash.String(),
		Nonce:        nonce,
		StorageHash:  result.StorageHash.String(),
		AccountProof: result.AccountProof,
		StorageProof: storageProof,
	}
	proofBz, err := json.Marshal(proof)
	if err != nil {
		return nil, err
	}
	return proofBz, nil
}

func (eth *Eth) GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error {

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
		Difficulty:  blockRes.Difficulty().String(),
		Height: tibcclient.Height{
			RevisionNumber: 0,
			RevisionHeight: req.LatestHeight,
		},
		GasLimit:  blockRes.GasLimit(),
		GasUsed:   blockRes.GasUsed(),
		Time:      blockRes.Time(),
		Extra:     blockRes.Extra(),
		MixDigest: blockRes.MixDigest().Bytes(),
		Nonce:     blockRes.Nonce(),
		BaseFee:   blockRes.BaseFee().String(),
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

func (eth *Eth) GetResult(hash string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()

	cmnHash := gethcmn.HexToHash(hash)
	result, err := eth.ethClient.TransactionReceipt(ctx, cmnHash)
	if err != nil {
		return 0, err
	}
	return result.Status, nil
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

func (eth *Eth) getProof(ctx context.Context, account gethcmn.Address, keys []string, blockNumber *big.Int) (*gethclient.AccountResult, error) {
	type storageResult struct {
		Key   string       `json:"key"`
		Value *hexutil.Big `json:"value"`
		Proof []string     `json:"proof"`
	}

	type accountResult struct {
		Address      gethcmn.Address `json:"address"`
		AccountProof []string        `json:"accountProof"`
		Balance      *hexutil.Big    `json:"balance"`
		CodeHash     gethcmn.Hash    `json:"codeHash"`
		Nonce        hexutil.Uint64  `json:"nonce"`
		StorageHash  gethcmn.Hash    `json:"storageHash"`
		StorageProof []storageResult `json:"storageProof"`
	}

	var res accountResult
	err := eth.gethRpcCli.CallContext(ctx, &res, "eth_getProof", account, keys, toBlockNumArg(blockNumber))

	// Turn hexutils back to normal datatypes
	storageResults := make([]gethclient.StorageResult, 0, len(res.StorageProof))
	for _, st := range res.StorageProof {
		storageResults = append(storageResults, gethclient.StorageResult{
			Key:   st.Key,
			Value: st.Value.ToInt(),
			Proof: st.Proof,
		})
	}
	result := &gethclient.AccountResult{
		Address:      res.Address,
		AccountProof: res.AccountProof,
		Balance:      res.Balance.ToInt(),
		Nonce:        uint64(res.Nonce),
		CodeHash:     res.CodeHash,
		StorageHash:  res.StorageHash,
		StorageProof: storageResults,
	}
	return result, err
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
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
func (eth *Eth) getCleanPacket(height uint64) ([]packet.CleanPacket, error) {
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

func (eth *Eth) getLogs(address gethcmn.Address, topic string, fromBlock, toBlock uint64) ([]gethtypes.Log, error) {
	filter := geth.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []gethcmn.Address{address},
		Topics:    [][]gethcmn.Hash{{gethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return eth.ethClient.FilterLogs(context.Background(), filter)
}

// ==================================================================================================================
// contract bind opts
type bindOpts struct {
	client             *bind.TransactOpts
	packetTransactOpts *bind.TransactOpts
}

func newBindOpts(cfg *ContractBindOptsCfg) (*bindOpts, error) {

	cliPriv, err := gethcrypto.HexToECDSA(cfg.ClientPrivKey)
	if err != nil {
		return nil, err
	}
	clientOpts, err := bind.NewKeyedTransactorWithChainID(cliPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	clientOpts.GasLimit = cfg.GasLimit
	clientOpts.GasPrice = new(big.Int).SetUint64(cfg.GasPrice)

	//================================================================================
	// packet transfer opts
	packPriv, err := gethcrypto.HexToECDSA(cfg.PacketPrivKey)
	if err != nil {
		return nil, err
	}
	packOpts, err := bind.NewKeyedTransactorWithChainID(packPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	packOpts.GasLimit = cfg.GasLimit
	packOpts.GasPrice = new(big.Int).SetUint64(cfg.GasPrice)

	return &bindOpts{
		client:             clientOpts,
		packetTransactOpts: packOpts,
	}, nil
}

// ==================================================================================================================
// contract client group
type contractGroup struct {
	Packet *contracts.Packet
	Client *contracts.Client
}

func newContractGroup(ethClient *gethethclient.Client, cfgGroup *ContractCfgGroup) (*contractGroup, error) {
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

	return &contractGroup{
		Packet: packetFilter,
		Client: clientFilter,
	}, nil
}
