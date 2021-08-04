package repostitory

import (
	"context"
	"fmt"
	sdk "github.com/irisnet/core-sdk-go"
	"github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/store"
)

type TendermintClient struct {
	delay     uint64
	chainName string
	sdk.Client
	rootAccount  Account
	randAccounts []Account
}
type Account struct {
	Name, Password string
	Address        types.AccAddress
}
type Config struct {
	NodeURI  string
	GrpcAddr string
	ChainID  string
}

func NewTendermintClient(chaiName string) *TendermintClient {
	config := Config{
		"tcp://127.0.0.1:26657",
		"localhost:9090",
		"testnet",
	}
	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
	}
	cfg, err := types.NewClientConfig(config.NodeURI, config.GrpcAddr, config.ChainID, options...)
	if err != nil {
		panic(err)
	}
	client := sdk.NewClient(cfg)
	return &TendermintClient{
		delay:     0,
		chainName: chaiName,
		Client:    client,
	}
}

func (c *TendermintClient) GetBlockAndPackets(height uint64) (interface{}, error) {
	a := int64(height)
	block, err := c.Client.Block(context.Background(), &a)
	if err != nil {
		panic(err)
	}
	c.Status(context.Background())
	return block, nil
}

func (c TendermintClient) GetBlockHeader(height uint64) (header interface{}, errMsg error) {
	tmp := int64(height)
	block, err := c.Client.Block(context.Background(), &tmp)
	if err != nil {
		errMsg = err
	}
	header = block.Block.Header
	return header, errMsg
}
func (c TendermintClient) GetLightClientState(chainName string) (interface{}, error) {
	res, err := c.Client.Status(context.Background())
	return res, err
}
func (c TendermintClient) GetLightClientConsensusState(chainName string, height uint64) (interface{}, error) {
	var tmp = int64(height)
	res, err := c.Client.ConsensusParams(context.Background(), &tmp)
	fmt.Println(res.BlockHeight)
	return res, err
}
func (c TendermintClient) GetStatus() (interface{}, error) {
	status, err := c.Client.Status(context.Background())
	return status, err
}
func (c TendermintClient) GetLatestHeight() (uint64, error) {
	block, err := c.Client.Block(context.Background(), nil)
	height := block.Block.Height
	return uint64(height), err
}
func (c TendermintClient) GetDelay() uint64 {
	//c.Client.Block()
	return c.delay
}
