package repostitory

import (
	"context"
	"math/big"
	"time"

	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	geth "github.com/ethereum/go-ethereum"
	gethcmn "github.com/ethereum/go-ethereum/common"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/irisnet/core-sdk-go/types"
)

var _ IChain = new(Eth)

const CtxTimeout = 10 * time.Second

type Eth struct {
	uri                   string
	chainName             string
	chainType             string
	updateClientFrequency uint64

	ethClient *gethclient.Client
}

func NewEth(chainType, chainName string, updateClientFrequency uint64, uri string) (IChain, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	ethClient, err := gethclient.DialContext(ctx, uri)
	if err != nil {
		return nil, err
	}

	return &Eth{
		chainType:             chainType,
		chainName:             chainName,
		updateClientFrequency: updateClientFrequency,
		ethClient:             ethClient,
	}, nil
}

func (eth *Eth) GetPackets(height uint64) (*repotypes.Packets, error) {
	block, err := eth.ethClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}
	blockHash := block.Hash()
	filter := geth.FilterQuery{
		BlockHash: &blockHash,
		Addresses: []gethcmn.Address{gethcmn.HexToAddress("")},
		Topics:    [][]gethcmn.Hash{{gethcrypto.Keccak256Hash([]byte(""))}},
	}
	eth.ethClient.FilterLogs(context.Background(), filter)

	return nil, nil
}

func (eth *Eth) GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error) {
	return nil, nil
}

func (eth *Eth) RecvPackets(msgs types.Msgs) (types.ResultTx, types.Error) {
	return types.ResultTx{}, nil
}

func (eth *Eth) GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) (*packet.QueryPacketCommitmentResponse, error) {
	return nil, nil
}

func (eth *Eth) GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (*packet.QueryPacketReceiptResponse, error) {
	return nil, nil
}

func (eth *Eth) GetBlockHeader(*repotypes.GetBlockHeaderReq) (tibctypes.Header, error) {
	return nil, nil
}

func (eth *Eth) GetLightClientState(string) (tibctypes.ClientState, error) {
	return nil, nil
}

func (eth *Eth) GetLightClientConsensusState(string, uint64) (tibctypes.ConsensusState, error) {
	return nil, nil
}

func (eth *Eth) GetStatus() (interface{}, error) {
	return nil, nil
}

func (eth *Eth) GetLatestHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return eth.ethClient.BlockNumber(ctx)
}

func (eth *Eth) GetLightClientDelayHeight(string) (uint64, error) {
	return 0, nil
}

func (eth *Eth) GetLightClientDelayTime(string) (uint64, error) {
	return 0, nil
}

func (eth *Eth) UpdateClient(header tibctypes.Header, chainName string) (string, error) {
	return "", nil
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

type Contracts struct {
	Client Contract
	Packet Contract
}

type Contract struct {
	Addr  string
	topic string
}

func NewContracts() *Contracts {
	return &Contracts{}
}
