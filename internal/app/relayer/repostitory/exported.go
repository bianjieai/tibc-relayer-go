package repostitory

import (
	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types"
)

type IChain interface {
	GetPackets(height uint64, destChainType string) (*repotypes.Packets, error)
	GetProof(sourChainName, destChainName string, sequence uint64, height uint64, typ string) ([]byte, error)
	RecvPackets(msgs types.Msgs) (*repotypes.ResultTx, types.Error)
	GetCommitmentsPacket(sourChainName, destChainName string, sequence uint64) error
	GetReceiptPacket(sourChainName, destChianName string, sequence uint64) (bool, error)
	GetBlockHeader(*repotypes.GetBlockHeaderReq) (tibctypes.Header, error)
	GetBlockTimestamp(height uint64) (uint64, error)
	GetLightClientState(string) (tibctypes.ClientState, error)
	GetLightClientConsensusState(string, uint64) (tibctypes.ConsensusState, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight(string) (uint64, error)
	GetLightClientDelayTime(string) (uint64, error)
	UpdateClient(header tibctypes.Header, chainName string) (string, error)

	GetResult(hash string) (uint64, error)

	ChainName() string
	UpdateClientFrequency() uint64
	ChainType() string
}
