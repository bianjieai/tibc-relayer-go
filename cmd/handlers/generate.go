package handlers

import (
	"bufio"
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	tibceth "github.com/bianjieai/tibc-sdk-go/eth"

	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common/hexutil"

	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	coresdk "github.com/irisnet/core-sdk-go"
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
)

const TendermintAndTendermint = "tendermint_and_tendermint"
const TendermintAndETH = "tendermint_and_eth"

const tibcTendermintMerklePrefix = "tibc"
const tibcTendermintRoot = "app_hash"

const (
	clientStatePrefix = `{"@type":"/tibc.lightclients.tendermint.v1.ClientState",`

	consensusStatePrefix = `{"@type":"/tibc.lightclients.tendermint.v1.ConsensusState",`

	EthConsensusStatePrefix = `{"@type":"/tibc.lightclients.eth.v1.ConsensusState",`
	EthClientStatePrefix    = `{"@type":"/tibc.lightclients.eth.v1.ClientState",`
)

func CreateClientFiles(cfg *configs.Config) {

	for _, channelType := range cfg.App.ChannelTypes {
		switch channelType {
		case TendermintAndTendermint:
			logger := log.WithFields(log.Fields{
				"source_chain": &cfg.Chain.Source.Tendermint.ChainName,
				"dest_chain":   &cfg.Chain.Dest.Tendermint.ChainName,
			})

			logger.Info("1. init source chain")
			sourceChain := tendermintCreateClientFiles(&cfg.Chain.Source, logger)
			getTendermintJson(
				sourceChain,
				int64(cfg.Chain.Source.Cache.StartHeight),
				cfg.Chain.Source.Tendermint.ChainName,
			)

			logger.Info("2. init dest chain")
			destChain := tendermintCreateClientFiles(&cfg.Chain.Dest, logger)
			getTendermintJson(
				destChain,
				int64(cfg.Chain.Dest.Cache.StartHeight),
				cfg.Chain.Dest.Tendermint.ChainName,
			)
		case TendermintAndETH:
			logger := log.WithFields(log.Fields{
				"source_chain": &cfg.Chain.Source.Tendermint.ChainName,
				"dest_chain":   &cfg.Chain.Dest.Eth.ChainName,
			})
			logger.Info("1. init source chain")
			sourceChain := tendermintCreateClientFiles(&cfg.Chain.Source, logger)
			getTendermintBytes(
				sourceChain,
				int64(cfg.Chain.Source.Cache.StartHeight),
				cfg.Chain.Source.Tendermint.ChainName,
			)
			logger.Info("2. init dest chain")
			getETHJson(&cfg.Chain.Dest, sourceChain, logger)
		}
	}
}

func getETHJson(cfg *configs.ChainCfg, client coresdk.Client, logger *log.Entry) {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, cfg.Eth.URI)
	if err != nil {
		logger.Fatal(err)
	}
	ethClient := gethethclient.NewClient(rpcClient)
	latestHeight, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	startHeight := latestHeight - 100
	logger.Info("eth height = ", startHeight)

	//gethCli := gethclient.New(rpcClient)
	blockRes, err := ethClient.BlockByNumber(
		context.Background(),
		new(big.Int).SetUint64(startHeight))
	if err != nil {
		logger.Fatal(err)
	}

	blockHeader := blockRes.Header()
	header := &tibceth.EthHeader{
		ParentHash:  blockHeader.ParentHash,
		UncleHash:   blockHeader.UncleHash,
		Coinbase:    blockHeader.Coinbase,
		Root:        blockHeader.Root,
		TxHash:      blockHeader.TxHash,
		ReceiptHash: blockHeader.ReceiptHash,
		Bloom:       blockHeader.Bloom,
		Difficulty:  blockHeader.Difficulty,
		Number:      blockHeader.Number,
		GasLimit:    blockHeader.GasLimit,
		GasUsed:     blockHeader.GasUsed,
		Time:        blockHeader.Time,
		Extra:       blockHeader.Extra,
		MixDigest:   blockHeader.MixDigest,
		Nonce:       blockHeader.Nonce,
		BaseFee:     blockHeader.BaseFee,
	}
	number := tibcclient.NewHeight(0, header.Number.Uint64())

	clientState := &tibceth.ClientState{
		Header:          header.ToHeader(),
		ChainId:         cfg.Eth.ChainID,
		ContractAddress: []byte(cfg.Eth.Contracts.Client.Addr),
		TrustingPeriod:  200,
		TimeDelay:       0,
		BlockDelay:      7,
	}
	consensusState := &tibceth.ConsensusState{
		Timestamp: header.Time,
		Number:    number,
		Root:      header.Root[:],
		Header:    header.ToHeader(),
	}

	clientStateBytes, err := client.AppCodec().MarshalJSON(clientState)
	if err != nil {
		logger.Fatal(err)
	}

	clientStateStr := string(clientStateBytes)
	clientStateStr = EthClientStatePrefix + clientStateStr[1:]
	clientStateFilename := fmt.Sprintf("%s_clientState.json", cfg.Eth.ChainName)
	writeCreateClientFiles(clientStateFilename, clientStateStr)

	consensusStateBytes, err := client.AppCodec().MarshalJSON(consensusState)
	if err != nil {
		logger.Fatal(err)
	}

	consensusStateStr := string(consensusStateBytes)
	consensusStateStr = EthConsensusStatePrefix + consensusStateStr[1:]
	consensusStateFilename1 := fmt.Sprintf("%s_consensusState.json", cfg.Eth.ChainName)
	writeCreateClientFiles(consensusStateFilename1, consensusStateStr)
}

func tendermintCreateClientFiles(cfg *configs.ChainCfg, logger *log.Entry) coresdk.Client {
	chainCfg := repostitory.NewTerndermintConfig()
	chainCfg.ChainID = cfg.Tendermint.ChainID
	chainCfg.GrpcAddr = cfg.Tendermint.GrpcAddr
	chainCfg.RPCAddr = cfg.Tendermint.RPCAddr

	fee := coretypes.NewDecCoins(
		coretypes.NewDecCoin(
			cfg.Tendermint.Fee.Denom,
			coretypes.NewInt(cfg.Tendermint.Fee.Amount)))

	chainCfg.BaseTx = coretypes.BaseTx{
		From:               cfg.Tendermint.Key.Name,
		Password:           cfg.Tendermint.Key.Password,
		Gas:                cfg.Tendermint.Gas,
		Mode:               coretypes.Commit,
		Fee:                fee,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}
	chainCfg.Name = cfg.Tendermint.Key.Name
	chainCfg.Password = cfg.Tendermint.Key.Password
	chainCfg.PrivKeyArmor = cfg.Tendermint.Key.PrivKeyArmor
	chainCfg.Options = []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(corestore.NewMemory(nil))),
		coretypes.TimeoutOption(10),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(cfg.Tendermint.Gas),
		coretypes.CachedOption(true),
	}

	coreSdkCfg, err := coretypes.NewClientConfig(
		chainCfg.RPCAddr, chainCfg.GrpcAddr, chainCfg.ChainID, chainCfg.Options...)
	if err != nil {
		logger.Fatal(err)
	}

	return coresdk.NewClient(coreSdkCfg)
}

func getTendermintBytes(client coresdk.Client, height int64, chainName string) {

	//ClientState
	var fra = tendermint.Fraction{
		Numerator:   1,
		Denominator: 3,
	}
	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header
	lastHeight := tibcclient.NewHeight(0, 4)
	var clientState = &tendermint.ClientState{
		ChainId:         tmHeader.ChainID,
		TrustLevel:      fra,
		TrustingPeriod:  time.Hour * 24 * 7 * 2,
		UnbondingPeriod: time.Hour * 24 * 7 * 3,
		MaxClockDrift:   time.Second * 10,
		LatestHeight:    lastHeight,
		ProofSpecs:      commitment.GetSDKSpecs(),
		MerklePrefix:    commitment.MerklePrefix{KeyPrefix: []byte(tibcTendermintMerklePrefix)},
		TimeDelay:       0,
	}
	//ConsensusState
	var consensusState = &tendermint.ConsensusState{
		Timestamp:          tmHeader.Time,
		Root:               commitment.NewMerkleRoot([]byte(tibcTendermintRoot)),
		NextValidatorsHash: tendermintQueryValidatorSet(res.Block.Height, client).Hash(),
	}

	clientStateBytes, err := clientState.Marshal()
	if err != nil {
		log.Fatal(err)
	}

	clientStateBytes1, err := client.AppCodec().MarshalJSON(clientState)
	if err != nil {
		panic(err)
	}
	// write file
	clientStateStr1 := string(clientStateBytes1)
	clientStateStr1 = clientStatePrefix + clientStateStr1[1:]
	clientStateFilename1 := fmt.Sprintf("%s_clientState.json", chainName)
	writeCreateClientFiles(clientStateFilename1, clientStateStr1)

	consensusStateBytes1, err := client.AppCodec().MarshalJSON(consensusState)
	if err != nil {
		panic(err)
	}
	consensusStateStr1 := string(consensusStateBytes1)
	consensusStateStr1 = consensusStatePrefix + consensusStateStr1[1:]
	consensusStateFilename1 := fmt.Sprintf("%s_consensusState.json", chainName)
	writeCreateClientFiles(consensusStateFilename1, consensusStateStr1)

	// write file
	clientStateFilename := fmt.Sprintf("%s_lientState.txt", chainName)
	writeCreateClientFiles(clientStateFilename, hexutil.Encode(clientStateBytes))

	consensusStateBytes, err := consensusState.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	consensusStateFilename := fmt.Sprintf("%s_consensusState.txt", chainName)
	writeCreateClientFiles(consensusStateFilename, hexutil.Encode(consensusStateBytes))
}

func getTendermintJson(client coresdk.Client, height int64, chainName string) {

	//ClientState
	var fra = tendermint.Fraction{
		Numerator:   1,
		Denominator: 3,
	}
	res, err := client.QueryBlock(height)
	if err != nil {
		fmt.Println("QueryBlock fail:  ", err)
	}
	tmHeader := res.Block.Header
	fmt.Println(tmHeader.ChainID)
	lastHeight := tibcclient.NewHeight(0, 4)
	var clientState = &tendermint.ClientState{
		ChainId:         tmHeader.ChainID,
		TrustLevel:      fra,
		TrustingPeriod:  time.Hour * 24 * 7 * 2,
		UnbondingPeriod: time.Hour * 24 * 7 * 3,
		MaxClockDrift:   time.Second * 10,
		LatestHeight:    lastHeight,
		ProofSpecs:      commitment.GetSDKSpecs(),
		MerklePrefix:    commitment.MerklePrefix{KeyPrefix: []byte(tibcTendermintMerklePrefix)},
		TimeDelay:       0,
	}
	//ConsensusState
	var consensusState = &tendermint.ConsensusState{
		Timestamp:          tmHeader.Time,
		Root:               commitment.NewMerkleRoot([]byte(tibcTendermintRoot)),
		NextValidatorsHash: tendermintQueryValidatorSet(res.Block.Height, client).Hash(),
	}

	clientStateBytes, err := client.AppCodec().MarshalJSON(clientState)
	if err != nil {
		panic(err)
	}
	// write file
	clientStateStr := string(clientStateBytes)
	clientStateStr = clientStatePrefix + clientStateStr[1:]
	clientStateFilename := fmt.Sprintf("%s_clientState.json", chainName)
	writeCreateClientFiles(clientStateFilename, clientStateStr)

	consensusStateBytes, err := client.AppCodec().MarshalJSON(consensusState)
	if err != nil {
		panic(err)
	}
	consensusStateStr := string(consensusStateBytes)
	consensusStateStr = consensusStatePrefix + consensusStateStr[1:]
	consensusStateFilename := fmt.Sprintf("%s_consensusState.json", chainName)
	writeCreateClientFiles(consensusStateFilename, consensusStateStr)
}

func writeCreateClientFiles(filePath string, content string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(content); err != nil {
		panic(err)
	}
	writer.Flush()
}

func tendermintQueryValidatorSet(height int64, client coresdk.Client) *tmtypes.ValidatorSet {
	validators, err := client.Validators(context.Background(), &height, nil, nil)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	validatorSet := tmtypes.NewValidatorSet(validators.Validators)
	if err != nil {
		fmt.Println("queryValidatorSet fail :", err)
	}
	return validatorSet
}
