package repostitory

type IChain interface {
	GetBlockAndPackets(height uint64) (interface{}, error)
	GetBlockHeader(height uint64) (interface{}, error)
	GetLightClientState(chainName string) (interface{}, error)
	GetLightClientConsensusState(chainName string, height uint64) (interface{}, error)
	GetStatus() (interface{}, error)
	GetLatestHeight() (uint64, error)
	GetDelay() uint64
}
