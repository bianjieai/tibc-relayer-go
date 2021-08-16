package repostitory

import (
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

type IChain interface {
	GetBlockAndPackets(height uint64) (interface{}, error)
	GetBlockHeader(height uint64) (tibctypes.Header, error)
	GetLightClientState(chainName string) (tibctypes.ClientState, error)
	GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error)
	GetStatus() (interface{}, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight() uint64
	GetLightClientDelayTime() uint64
	UpdateClient(header tibctypes.Header) error
	ChainName() string
}
