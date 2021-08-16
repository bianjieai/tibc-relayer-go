package repostitory

import (
	"context"

	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	"github.com/gogo/protobuf/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (c *TendermintClient) GetBlockAndPackets(height uint64) (interface{}, error) {
	a := int64(height)
	return c.BaseClient.Block(context.Background(), &a)
}

func (c *TendermintClient) GetBlockHeader(height uint64, trustedHeight tibcclient.Height, trustedValidators *tenderminttypes.ValidatorSet) (tibctypes.Header, error) {
	block, err := c.QueryBlock(int64(height))
	if err != nil {
		return nil, err
	}
	rescommit, err := c.Commit(context.Background(), &block.BlockResult.Height)
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

func (c *TendermintClient) GetLightClientState(chainName string) (tibctypes.ClientState, error) {
	return c.TendermintClient.GetClientState(chainName)

}

func (c *TendermintClient) GetLightClientConsensusState(chainName string, height uint64) (tibctypes.ConsensusState, error) {
	return c.TendermintClient.GetConsensusState(chainName, height)
}

func (c *TendermintClient) GetStatus() (interface{}, error) {
	return c.BaseClient.Status(context.Background())
}

func (c *TendermintClient) GetLatestHeight() (uint64, error) {
	block, err := c.BaseClient.Block(context.Background(), nil)
	var height = block.Block.Height
	return uint64(height), err
}

func (c *TendermintClient) GetLightClientDelayHeight(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	return res.GetDelayBlock(), err
}

func (c *TendermintClient) GetLightClientDelayTime(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	return res.GetDelayTime(), err
}

func (c *TendermintClient) UpdateClient(header tibctypes.Header, chainName string, baseTx coretypes.BaseTx) error {
	request := tibctypes.UpdateClientRequest{
		ChainName: "testCreateClient1",
		Header:    header,
	}
	_, err := c.TendermintClient.UpdateClient(request, baseTx)
	if err != nil {
		return err
	}
	return nil
}

func (c *TendermintClient) ChainName() string {
	return c.chainName
}

func (c *TendermintClient) queryValidatorSet(height int64) (*tenderminttypes.ValidatorSet, error) {

	validators, err := c.Validators(context.Background(), &height, nil, nil)
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
