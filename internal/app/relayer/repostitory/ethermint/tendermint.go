package ethermint

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"regexp"
	"time"

	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/tendermint/tendermint/light/provider"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/irisnet/core-sdk-go/bank"
	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/gov"
	"github.com/irisnet/core-sdk-go/staking"
	"github.com/irisnet/irismod-sdk-go/nft"

	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcmttypes "github.com/bianjieai/tibc-sdk-go/modules/apps/mt_transfer"
	tibcclient "github.com/bianjieai/tibc-sdk-go/modules/core/client"
	"github.com/bianjieai/tibc-sdk-go/modules/light-clients/tendermint"
	tibctypes "github.com/bianjieai/tibc-sdk-go/modules/types"
	"github.com/irisnet/core-sdk-go/common/codec"
	cdctypes "github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	"github.com/irisnet/core-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"
	"github.com/tendermint/tendermint/libs/log"
	tenderminttypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	repotypes "github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory/types"
)

var _ repostitory.IChain = new(Ethermint)

var (
	maxRetryAttempts    = 5
	regexpTooHigh       = regexp.MustCompile(`height \d+ must be less than or equal to`)
	regexpMissingHeight = regexp.MustCompile(`height \d+ is not available`)
	regexpTimedOut      = regexp.MustCompile(`Timeout exceeded`)
)

type Ethermint struct {
	logger log.Logger

	chainName             string
	chainType             string
	updateClientFrequency uint64

	// tendermint config
	terndermintCli tendermintClient

	// eth config
	uri              string
	contractCfgGroup *ContractCfgGroup
	contracts        *contractGroup
	bindOpts         *bindOpts

	slot           int64
	maxGasPrice    *big.Int
	tipCoefficient float64

	ethClient  *gethethclient.Client
	gethCli    *gethclient.Client
	gethRpcCli *gethrpc.Client
}

func NewEthermintClient(
	chainType string,
	chainName string,
	updateClientFrequency uint64,
	config *Config) (*Ethermint, error) {

	// init eth client
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, config.Eth.ChainURI)
	if err != nil {
		return nil, err
	}

	ethClient := gethethclient.NewClient(rpcClient)
	gethCli := gethclient.New(rpcClient)

	contractGroup, err := newContractGroup(ethClient, config.Eth.ContractCfgGroup)
	if err != nil {
		return nil, err
	}

	tmpBindOpts, err := newBindOpts(config.Eth.ContractBindOptsCfg)

	if err != nil {
		return nil, err
	}

	// init tendermint client
	cfg, err := coretypes.NewClientConfig(
		config.Tendermint.RPCAddr,
		config.Tendermint.GrpcAddr,
		config.Tendermint.ChainID,
		config.Tendermint.Options...)
	if err != nil {
		return nil, err
	}
	tc := newTendermintClient(cfg, chainName)

	return &Ethermint{
		//Common properties
		chainType:             chainType,
		chainName:             chainName,
		updateClientFrequency: updateClientFrequency,

		terndermintCli: tc,
		logger:         tc.BaseClient.Logger(),

		// eth
		contractCfgGroup: config.Eth.ContractCfgGroup,
		ethClient:        ethClient,
		gethCli:          gethCli,
		gethRpcCli:       rpcClient,
		contracts:        contractGroup,
		bindOpts:         tmpBindOpts,
		slot:             config.Eth.Slot,
		tipCoefficient:   config.Eth.TipCoefficient,
		maxGasPrice:      new(big.Int).SetUint64(config.Eth.ContractBindOptsCfg.MaxGasPrice),
	}, err
}

func (c *Ethermint) GetBlockTimestamp(height uint64) (uint64, error) {
	block, err := c.terndermintCli.QueryBlock(int64(height))
	if err != nil {
		return 0, err
	}

	return uint64(block.Block.Time.Unix()), nil
}

func (c *Ethermint) GetBlockHeader(req *repotypes.GetBlockHeaderReq) (tibctypes.Header, error) {
	block, err := c.terndermintCli.QueryBlock(int64(req.LatestHeight))
	if err != nil {
		return nil, err
	}
	rescommit, err := c.terndermintCli.Commit(context.Background(), &block.BlockResult.Height)
	if err != nil {
		return nil, err
	}
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
			RevisionNumber: req.RevisionNumber,
			RevisionHeight: req.TrustedHeight,
		},
		TrustedValidators: trustedValidators,
	}, nil
}

func (c *Ethermint) GetLatestHeight() (uint64, error) {
	block, err := c.terndermintCli.Block(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	var height = block.Block.Height
	return uint64(height), err
}

func (c *Ethermint) GetResult(hash string) (uint64, error) {
	res, err := c.terndermintCli.QueryTx(hash)
	if err != nil {
		return 0, err
	}
	code := uint64(res.Result.Code)
	return code, nil
}

func (c *Ethermint) GetLightClientDelayHeight(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayBlock(), nil
}

func (c *Ethermint) GetLightClientDelayTime(chainName string) (uint64, error) {
	res, err := c.GetLightClientState(chainName)
	if err != nil {
		return 0, err
	}
	return res.GetDelayTime(), nil

}

func (c *Ethermint) ChainName() string {

	return c.chainName
}

func (c *Ethermint) ChainType() string {
	return c.chainType
}

func (c *Ethermint) UpdateClientFrequency() uint64 {
	return c.updateClientFrequency
}

func (c *Ethermint) getValidator(height int64) (*tenderminttypes.ValidatorSet, error) {
	const maxPages = 100

	var (
		perPage = 100
		vals    = []*tmtypes.Validator{}
		page    = 1
		total   = -1
	)
	ctx := context.Background()

OUTER_LOOP:
	for len(vals) != total && page <= maxPages {
		for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
			res, err := c.terndermintCli.TIBC.Validators(ctx, &height, &page, &perPage)
			switch {
			case err == nil:
				// Validate response.
				if len(res.Validators) == 0 {
					return nil, provider.ErrBadLightBlock{
						Reason: fmt.Errorf("validator set is empty (height: %d, page: %d, per_page: %d)",
							height, page, perPage),
					}
				}
				if res.Total <= 0 {
					return nil, provider.ErrBadLightBlock{
						Reason: fmt.Errorf("total number of vals is <= 0: %d (height: %d, page: %d, per_page: %d)",
							res.Total, height, page, perPage),
					}
				}

				total = res.Total
				vals = append(vals, res.Validators...)
				page++
				continue OUTER_LOOP

			case regexpTooHigh.MatchString(err.Error()):
				return nil, fmt.Errorf("height requested is too high")

			case regexpMissingHeight.MatchString(err.Error()):
				return nil, provider.ErrLightBlockNotFound

			// if we have exceeded retry attempts then return no response error
			case attempt == maxRetryAttempts:
				return nil, provider.ErrNoResponse

			case regexpTimedOut.MatchString(err.Error()):
				// we wait and try again with exponential backoff
				time.Sleep(backoffTimeout(uint16(attempt)))
				continue

			// context canceled or connection refused we return the error
			default:
				return nil, err
			}

		}
	}
	validatorSet, err := tmtypes.NewValidatorSet(vals).ToProto()
	if err != nil {
		return nil, err
	}

	return validatorSet, nil
}

// exponential backoff (with jitter)
// 0.5s -> 2s -> 4.5s -> 8s -> 12.5 with 1s variation
func backoffTimeout(attempt uint16) time.Duration {
	// nolint:gosec // G404: Use of weak random number generator
	return time.Duration(500*attempt*attempt)*time.Millisecond + time.Duration(rand.Intn(1000))*time.Millisecond
}

//======================================

type tendermintClient struct {
	encodingConfig types.EncodingConfig
	coretypes.BaseClient
	Bank      bank.Client
	Staking   staking.Client
	Gov       gov.Client
	NFT       nft.Client
	TIBC      tibc.Client
	ChainName string
}

func newTendermintClient(cfg types.ClientConfig, chainName string) tendermintClient {
	encodingConfig := makeEncodingConfig()
	// create a instance of baseClient
	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	stakingClient := staking.NewClient(baseClient, encodingConfig.Marshaler)
	govClient := gov.NewClient(baseClient, encodingConfig.Marshaler)
	tibcClient := tibc.NewClient(baseClient, encodingConfig)
	nftClient := nft.NewClient(baseClient, encodingConfig.Marshaler)

	tc := &tendermintClient{
		encodingConfig: encodingConfig,
		BaseClient:     baseClient,
		Bank:           bankClient,
		Staking:        stakingClient,
		Gov:            govClient,
		NFT:            nftClient,
		TIBC:           tibcClient,
		ChainName:      chainName,
	}

	tc.RegisterModule(
		bankClient,
		stakingClient,
		govClient,
	)
	return *tc
}

func (tc tendermintClient) Manager() types.BaseClient {
	return tc.BaseClient
}

func (tc tendermintClient) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(tc.encodingConfig.InterfaceRegistry)
	}
}

//client init
func makeEncodingConfig() types.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	registerLegacyAminoCodec(encodingConfig.Amino)
	registerInterfaces(encodingConfig.InterfaceRegistry)
	tibcmttypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func registerLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func registerInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}
