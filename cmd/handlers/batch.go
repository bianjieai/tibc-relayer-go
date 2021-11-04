package handlers

import (
	"context"
	"math/big"
	"time"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	tibc "github.com/bianjieai/tibc-sdk-go"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	tibceth "github.com/bianjieai/tibc-sdk-go/eth"
	tibctypes "github.com/bianjieai/tibc-sdk-go/types"
	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/irisnet/core-sdk-go/bank"
	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/common/codec"
	cdctypes "github.com/irisnet/core-sdk-go/common/codec/types"
	cryptotypes "github.com/irisnet/core-sdk-go/common/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	"github.com/irisnet/core-sdk-go/gov"
	"github.com/irisnet/core-sdk-go/staking"
	"github.com/irisnet/core-sdk-go/types"
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"
	"github.com/irisnet/irismod-sdk-go/nft"
	log "github.com/sirupsen/logrus"
)

const CtxTimeout = 10 * time.Second

func BatchUpdateETHClient(cfg *configs.Config, endHeight uint64) {
	logger := log.WithFields(log.Fields{
		"source_chain": cfg.Chain.Source.Tendermint.ChainName,
		"dest_chain":   cfg.Chain.Dest.Eth.ChainName,
	})
	if len(cfg.App.ChannelTypes) != 1 {
		logger.Fatal("channel_types length must be 1")
	}
	for _, channelType := range cfg.App.ChannelTypes {
		if channelType != TendermintAndETH {
			logger.Fatal("only applicable for eth and tendermint")
		}
		batchUpdateETHClient(cfg, endHeight, logger)
	}

}

func batchUpdateETHClient(cfg *configs.Config, endHeight uint64, logger *log.Entry) {

	options := []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(corestore.NewMemory(nil))),
		coretypes.TimeoutOption(cfg.Chain.Source.Tendermint.RequestTimeout),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(cfg.Chain.Source.Tendermint.Gas),
		coretypes.CachedOption(true),
	}
	if cfg.Chain.Source.Tendermint.Algo != "" {
		options = append(options, coretypes.AlgoOption(cfg.Chain.Source.Tendermint.Algo))
	}
	chainCfg, err := coretypes.NewClientConfig(
		cfg.Chain.Source.Tendermint.RPCAddr,
		cfg.Chain.Source.Tendermint.GrpcAddr,
		cfg.Chain.Source.Tendermint.ChainID,
		options...,
	)
	if err != nil {
		logger.WithField("err_msg", err).Fatal("failed to init chain cfg")
	}

	fee := coretypes.NewDecCoins(
		coretypes.NewDecCoin(
			cfg.Chain.Source.Tendermint.Fee.Denom,
			coretypes.NewInt(cfg.Chain.Source.Tendermint.Fee.Amount)))
	baseTx := coretypes.BaseTx{
		From:               cfg.Chain.Source.Tendermint.Key.Name,
		Password:           cfg.Chain.Source.Tendermint.Key.Password,
		Gas:                cfg.Chain.Source.Tendermint.Gas,
		Mode:               coretypes.Async,
		Fee:                fee,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	tClient := newTendermintClient(chainCfg, cfg.Chain.Source.Tendermint.ChainName)
	address, err := tClient.BaseClient.Import(
		cfg.Chain.Source.Tendermint.Key.Name,
		cfg.Chain.Source.Tendermint.Key.Password,
		cfg.Chain.Source.Tendermint.Key.PrivKeyArmor)
	if err != nil {
		logger.WithField("err_msg", err).Fatal("failed to get owner err")
	}
	logger.WithField("address", address).Info("address output")

	owner, err := tClient.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		logger.WithField("err_msg", err).Fatal("failed to get owner err")
	}
	logger.WithField("owner", owner.String()).Info("owner output")
	ethCli, err := newEthClient(cfg.Chain.Dest.Eth.URI)
	if err != nil {
		logger.WithField("err_msg", err).Fatal("failed to eth client")
	}

	clientStatus, err := tClient.TIBC.GetClientState(cfg.Chain.Dest.Eth.ChainName)
	if err != nil {
		logger.WithField("err_msg", err).Fatal("failed to get eth light client status")
	}

	startHeight := clientStatus.GetLatestHeight().GetRevisionHeight() + 1
	logger.WithFields(log.Fields{
		"latest_height": clientStatus.GetLatestHeight().GetRevisionHeight(),
		"start_height":  startHeight,
	}).Info()

	var msgs []types.Msg
	for i := startHeight; i <= endHeight; i++ {
		header, err := getETHHeader(ethCli, i)
		if err != nil {
			logger.WithField("err_msg", err).Fatal("failed to eth client")
		}
		anyHeader, err := cryptotypes.NewAnyWithValue(header)
		if err != nil {
			logger.WithField("err_msg", err).Fatal("failed to get owner err")
		}
		msg := &tibcclient.MsgUpdateClient{
			ChainName: cfg.Chain.Dest.Eth.ChainName,
			Header:    anyHeader,
			Signer:    owner.String(),
		}

		msgs = append(msgs, msg)
	}
	for {
		if len(msgs) == 0 {
			logger.Info("no data need to relay")
			break
		}

		var relayMsg []types.Msg
		msgLength := len(msgs)
		if msgLength > 2 {
			relayMsg = msgs[:2]
		} else {
			relayMsg = msgs[:msgLength]
		}

		resultTx, err := tClient.BuildAndSend(relayMsg, baseTx)
		if err != nil {
			logger.WithField("err_msg", err).Fatal("failed to tx")
		}
		logger.WithFields(log.Fields{
			"tx_height":  resultTx.Height,
			"tx_hash":    resultTx.Hash,
			"gas_wanted": resultTx.GasWanted,
			"gas_used":   resultTx.GasUsed,
		}).Info("success")
		msgs = msgs[len(relayMsg):]
		time.Sleep(time.Second * 5)
	}

}

func getETHHeader(ethCli *gethethclient.Client, height uint64) (tibctypes.Header, error) {

	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	blockRes, err := ethCli.BlockByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return nil, err
	}
	return &tibceth.Header{
		ParentHash:  blockRes.ParentHash().Bytes(),
		UncleHash:   blockRes.UncleHash().Bytes(),
		Coinbase:    blockRes.Coinbase().Bytes(),
		Root:        blockRes.Root().Bytes(),
		TxHash:      blockRes.TxHash().Bytes(),
		ReceiptHash: blockRes.ReceiptHash().Bytes(),
		Bloom:       blockRes.Bloom().Bytes(),
		Difficulty:  blockRes.Difficulty().String(),
		Height: tibcclient.Height{
			RevisionNumber: 0,
			RevisionHeight: height,
		},
		GasLimit:  blockRes.GasLimit(),
		GasUsed:   blockRes.GasUsed(),
		Time:      blockRes.Time(),
		Extra:     blockRes.Extra(),
		MixDigest: blockRes.MixDigest().Bytes(),
		Nonce:     blockRes.Nonce(),
		BaseFee:   blockRes.BaseFee().String(),
	}, nil
}

func newEthClient(uri string) (*gethethclient.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, uri)
	if err != nil {
		return nil, err
	}

	ethClient := gethethclient.NewClient(rpcClient)
	return ethClient, nil
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

type tendermintClient struct {
	encodingConfig types.EncodingConfig
	types.BaseClient
	Bank      bank.Client
	Staking   staking.Client
	Gov       gov.Client
	NFT       nft.Client
	TIBC      tibc.Client
	ChainName string
}

func (tc tendermintClient) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(tc.encodingConfig.InterfaceRegistry)
	}
}

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
