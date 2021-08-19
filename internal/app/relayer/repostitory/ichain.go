package repostitory

import (
	"github.com/bianjieai/tibc-sdk-go/packet"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
)

type IChain interface {
	GetPackets(height uint64) (*Packets, error)
	GetProof(chainName string, sequence uint64, height uint64) ([]byte, error)
	RecvPackets(msgs types.Msgs) (types.ResultTx, types.Error)
	GetCommitmentsPacket(chainName string, sequence uint64) (*packet.QueryPacketCommitmentResponse, error)
	GetReceiptPacket(chainName string, sequence uint64) (*packet.QueryPacketReceiptResponse, error)
	GetAckPacket(chainName string, sequence uint64) (*packet.QueryPacketAcknowledgementResponse, error)
	GetBlockHeader(*GetBlockHeaderReq) (tibctypes.Header, error)
	GetLightClientState(string) (tibctypes.ClientState, error)
	GetLightClientConsensusState(string, uint64) (tibctypes.ConsensusState, error)
	GetStatus() (interface{}, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight(string) (uint64, error)
	GetLightClientDelayTime(string) (uint64, error)
	UpdateClient(tibctypes.Header, string) error
	ChainName() string
	UpdateClientFrequency() uint64
	ChainType() string
}

type IPacketRepo interface {
}
