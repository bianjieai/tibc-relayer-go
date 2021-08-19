package repostitory

import (
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

type IChain interface {
	GetBlockAndPackets(height uint64) (interface{}, error)
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
