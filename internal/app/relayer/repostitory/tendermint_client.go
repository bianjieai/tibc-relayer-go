package repostitory

import (
	"context"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/gogo/protobuf/types"
	sdk "github.com/irisnet/core-sdk-go"
	coretypes "github.com/irisnet/core-sdk-go/types"
	"github.com/tendermint/tendermint/libs/log"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var _ IChain = new(Tendermint)

type Tendermint struct {
	logger log.Logger

	CoreSdk    sdk.Client
	TibcClient tibc.Client

	chainName string
}

func NewClient(cfg coretypes.ClientConfig,chainName string) Tendermint {
	coreClient := sdk.NewClient(cfg)
	tibcClient := tibc.NewClient(coreClient.BaseClient, coreClient.AppCodec())
	client := &Tendermint{
		logger:     coreClient.BaseClient.Logger(),
		CoreSdk:    coreClient,
		TibcClient: tibcClient,
		chainName:  chainName,
	}
	client.CoreSdk.RegisterModule(
		tibcClient,
	)
	return *client
}

func (c *Tendermint) GetBlockAndPackets(height uint64) (interface{}, error) {
	a := int64(height)
	return c.CoreSdk.Block(context.Background(), &a)
}

func (c *Tendermint) GetBlockHeader(height uint64, trustedHeight tibcclient.Height, trustedValidators *tenderminttypes.ValidatorSet) (tibctypes.Header, error) {
	block, err := c.CoreSdk.QueryBlock(int64(height))
	if err != nil {
		return nil, err
	}
	rescommit, err := c.CoreSdk.Commit(context.Background(), &block.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: block.Block.Header.ToProto(),
		Commit: commit.ToProto(),
	}
	validatorSet, err := c.queryValidatorSet(block.Block.Height)
	if err != nil {
		return nil, err
	}
	// The trusted fields may be nil. They may be filled before relaying messages to a client.
	// The relayer is responsible for querying client and injecting appropriate trusted fields.
	return &tendermint.Header{
		SignedHeader:      signedHeader,
		ValidatorSet:      validatorSet,
		TrustedHeight:     trustedHeight,
		TrustedValidators: trustedValidators,
	}, nil
}

func (c *Tendermint) GetLightClientState(chainName string) (tibctypes.ClientState, error) {
	return c.TibcClient.GetClientState(chainName)

}

func (c *Tendermint) GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {
	return c.TibcClient.GetConsensusState(chainName, height)
}

func (c *Tendermint) GetStatus() (interface{}, error) {
	return c.TibcClient.Status(context.Background())
}

func (c *Tendermint) GetLatestHeight() (uint64, error) {
	block, err := c.CoreSdk.Block(context.Background(), nil)
	var height = block.Block.Height
	return uint64(height), err
}

func (c *Tendermint) GetLightClientDelayHeight(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	return res.GetDelayBlock(), err
}

func (c *Tendermint) GetLightClientDelayTime(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	return res.GetDelayTime(), err
}

func (c *Tendermint) UpdateClient(header tibctypes.Header, chainName string, baseTx coretypes.BaseTx) error {
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    header,
	}
	_, err := c.TibcClient.UpdateClient(request, baseTx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Tendermint) ChainName() string {
	return c.chainName
}

func (c *Tendermint) queryValidatorSet(height int64) (*tenderminttypes.ValidatorSet, error) {

	validators, err := c.CoreSdk.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		return nil, err
	}
	validatorSet, err := tmtypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		return nil, err
	}
	return validatorSet, nil
}

type TerndermintConfig struct {
	Options []types.Option

	RPCAddr  string
	GrpcAddr string
	ChainID  string
}

func NewTerndermintConfig() *TerndermintConfig {
	return &TerndermintConfig{}
}
