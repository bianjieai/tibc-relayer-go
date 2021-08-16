package repostitory

import (
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
)

type IChain interface {
	GetBlockAndPackets(height uint64) (interface{}, error)
	GetBlockHeader(height uint64,trustedHeight tibcclient.Height,trustedValidators *tenderminttypes.ValidatorSet) (tibctypes.Header, error)
	GetLightClientState(chainName string) (tibctypes.ClientState, error)
	GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error)
	GetStatus() (interface{}, error)
	GetLatestHeight() (uint64, error)
	GetLightClientDelayHeight(chainName string) (uint64,error)
	GetLightClientDelayTime(chainName string) (uint64,error)
	UpdateClient(header tibctypes.Header,chainName string,baseTx coretypes.BaseTx) error
	ChainName() string
}
