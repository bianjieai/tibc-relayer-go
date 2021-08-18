package repostitory

import (
	"context"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go"
	"github.com/irisnet/core-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	"github.com/tendermint/tendermint/libs/log"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmttypes "github.com/tendermint/tendermint/types"
)

var _ IChain = new(Tendermint)

type Tendermint struct {
	logger log.Logger

	coreSdk    sdk.Client
	tibcClient tibc.Client
	baseTx     types.BaseTx

	chainName             string
	chainType             string
	updateClientFrequency uint64
}

func NewTendermintClient(chainType, chaiName string, updateClientFrequency uint64, config *TerndermintConfig) (*Tendermint, error) {
	cfg, err := coretypes.NewClientConfig(config.RPCAddr, config.GrpcAddr, config.ChainID, config.Options...)
	if err != nil {
		return nil, err
	}
	coreClient := sdk.NewClient(cfg)
	tibcClient := tibc.NewClient(coreClient)
	return &Tendermint{
		chainType:             chainType,
		chainName:             chaiName,
		updateClientFrequency: updateClientFrequency,
		logger:                coreClient.BaseClient.Logger(),
		coreSdk:               coreClient,
		tibcClient:            tibcClient,
		baseTx:                config.BaseTx,
	}, err
}

func (c *Tendermint) GetBlockAndPackets(height uint64) (interface{}, error) {
	a := int64(height)
	return c.coreSdk.Block(context.Background(), &a)
}

func (c *Tendermint) GetBlockHeader(req *GetBlockHeaderReq) (tibctypes.Header, error) {
	block, err := c.coreSdk.QueryBlock(int64(req.LatestHeight))
	if err != nil {
		return nil, err
	}
	rescommit, err := c.coreSdk.Commit(context.Background(), &block.BlockResult.Height)
	commit := rescommit.Commit
	signedHeader := &tenderminttypes.SignedHeader{
		Header: block.Block.Header.ToProto(),
		Commit: commit.ToProto(),
	}
	validatorSet, err := c.getValidator(int64(req.LatestHeight))
	if err != nil {
		return nil, err

	}
	trustedValidators, err := c.getValidator(int64(req.TrustedHeight))
	if err != nil {
		return nil, err
	}
	// The trusted fields may be nil. They may be filled before relaying messages to a client.
	// The relayer is responsible for querying client and injecting appropriate trusted fields.
	return &tendermint.Header{
		SignedHeader: signedHeader,
		ValidatorSet: validatorSet,
		TrustedHeight: tibcclient.Height{
			RevisionHeight: req.TrustedHeight,
		},
		TrustedValidators: trustedValidators,
	}, nil
}

func (c *Tendermint) GetLightClientState(chainName string) (tibctypes.ClientState, error) {
	return c.tibcClient.GetClientState(chainName)

}

func (c *Tendermint) GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {
	return c.tibcClient.GetConsensusState(chainName, height)

}

func (c *Tendermint) GetStatus() (interface{}, error) {
	return c.coreSdk.Status(context.Background())
}

func (c *Tendermint) GetLatestHeight() (uint64, error) {
	block, err := c.coreSdk.Block(context.Background(), nil)
	var height = block.Block.Height
	return uint64(height), err
}

func (c *Tendermint) GetLightClientDelayHeight(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayBlock(), nil
}

func (c *Tendermint) GetLightClientDelayTime(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayTime(), nil

}

func (c *Tendermint) UpdateClient(header tibctypes.Header, chainName string) error {
	request := tibctypes.UpdateClientRequest{
		ChainName: chainName,
		Header:    header,
	}
	_, err := c.tibcClient.UpdateClient(request, c.baseTx)
	if err != nil {
		return err
	}
	return nil
}

func (c *Tendermint) ChainName() string {

	return c.chainName
}

func (c *Tendermint) ChainType() string {
	return c.chainType
}

func (c *Tendermint) UpdateClientFrequency() uint64 {
	return c.updateClientFrequency
}

func (c *Tendermint) getValidator(height int64) (*tenderminttypes.ValidatorSet, error) {
	validators, err := c.coreSdk.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		return nil, err
	}
	validatorSet, err := tmttypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		return nil, err
	}

	return validatorSet, nil
}

type TerndermintConfig struct {
	Options []coretypes.Option
	BaseTx  types.BaseTx

	RPCAddr  string
	GrpcAddr string
	ChainID  string
}

func NewTerndermintConfig() *TerndermintConfig {
	return &TerndermintConfig{}
}
