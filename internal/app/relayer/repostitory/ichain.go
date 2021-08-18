package repostitory

import (
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
)

type IChain interface {
	GetBlockAndPackets(height uint64) (interface{}, error)
	GetBlockHeader(req *GetBlockHeaderReq) (tibctypes.Header, error)
	GetLightClientState(chainName string) (tibctypes.ClientState, error)
	GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error)
	GetStatus() (interface{}, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight() uint64
	GetLightClientDelayTime() uint64
	GetLightClientValidator(height int64, chainName string) (*QueryLightClientValidatorResp, error)
	UpdateClient(header tibctypes.Header) error

	ChainName() string
	UpdateClientFrequency() uint64
	ChainType() string
}
