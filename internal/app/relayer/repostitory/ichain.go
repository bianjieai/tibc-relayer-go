package repostitory

import (
	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
)

type IChain interface {
	GetPackets(height uint64) (*repotypes.Packets, error)
	GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error)
	RecvPackets(msgs types.Msgs) (types.ResultTx, types.Error)
	GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) (*packet.QueryPacketCommitmentResponse, error)
	GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (*packet.QueryPacketReceiptResponse, error)
	GetBlockHeader(*repotypes.GetBlockHeaderReq) (tibctypes.Header, error)
	GetLightClientState(string) (tibctypes.ClientState, error)
	GetLightClientConsensusState(string, uint64) (tibctypes.ConsensusState, error)
	GetStatus() (interface{}, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight(string) (uint64, error)
	GetLightClientDelayTime(string) (uint64, error)
	UpdateClient(header tibctypes.Header, chainName string) (string, error)
	ChainName() string
	UpdateClientFrequency() uint64
	ChainType() string
}

type IPacketRepo interface {
}
