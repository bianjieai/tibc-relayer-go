package repostitory

import (
	"context"

	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go"
	coretypes "github.com/irisnet/core-sdk-go/types"
)

var _ IChain = new(TendermintClient)

type TendermintClient struct {
	sdk.Client

	chainName             string
	chainType             string
	updateClientFrequency uint64
}

func NewTendermintClient(chainType, chaiName string, updateClientFrequency uint64, config *TerndermintConfig) (*TendermintClient, error) {
	cfg, err := coretypes.NewClientConfig(config.RPCAddr, config.GrpcAddr, config.ChainID, config.Options...)
	if err != nil {
		return nil, err
	}
	return &TendermintClient{
		chainType:             chainType,
		chainName:             chaiName,
		updateClientFrequency: updateClientFrequency,
		Client:                sdk.NewClient(cfg),
	}, err
}

func (c *TendermintClient) GetBlockAndPackets(height uint64) (interface{}, error) {
	a := int64(height)
	return c.Client.Block(context.Background(), &a)
}

func (c *TendermintClient) GetBlockHeader(req *GetBlockHeaderReq) (tibctypes.Header, error) {
	// todo
	// get block header
	return nil, nil
}

func (c *TendermintClient) GetLightClientState(chainName string) (tibctypes.ClientState, error) {

	// todo
	// c.Client.Status(context.Background())

	return nil, nil
}

func (c *TendermintClient) GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {

	//var tmp = int64(height)
	//c.Client.ConsensusParams(context.Background(), &tmp)
	return nil, nil
}

func (c *TendermintClient) GetStatus() (interface{}, error) {
	return c.Client.Status(context.Background())
}

func (c *TendermintClient) GetLatestHeight() (uint64, error) {
	block, err := c.Client.Block(context.Background(), nil)
	var height = block.Block.Height
	return uint64(height), err
}

func (c *TendermintClient) GetLightClientDelayHeight() uint64 {

	// todo
	return 0
}

func (c *TendermintClient) GetLightClientDelayTime() uint64 {

	// todo
	return 0
}

func (c *TendermintClient) UpdateClient(header tibctypes.Header) error {
	return nil
}

func (c *TendermintClient) ChainName() string {
	return c.chainName
}

func (c *TendermintClient) ChainType() string {
	return c.chainType
}

func (c *TendermintClient) UpdateClientFrequency() uint64 {
	return c.updateClientFrequency
}

func (c *TendermintClient) GetLightClientValidator(height int64, chainName string) (*QueryLightClientValidatorResp, error) {
	return nil, nil
}

type TerndermintConfig struct {
	Options []coretypes.Option

	RPCAddr  string
	GrpcAddr string
	ChainID  string
}

func NewTerndermintConfig() *TerndermintConfig {
	return &TerndermintConfig{}
}
