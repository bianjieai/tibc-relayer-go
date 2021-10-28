package handlers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/common"

	tibceth "github.com/bianjieai/tibc-sdk-go/eth"

	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

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

const Tendermint = "tendermint"
const ETH = "eth"

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

			if cfg.Chain.Source.ChainType == Tendermint && cfg.Chain.Dest.ChainType == ETH {
				fmt.Println(channelType, "sdddddd")
				logger := log.WithFields(log.Fields{
					"source_chain": &cfg.Chain.Source.Tendermint.ChainName,
					"dest_chain":   &cfg.Chain.Dest.Eth.ChainName,
				})
				logger.Info("1. init source chain")
				sourceChain := tendermintCreateClientFiles(&cfg.Chain.Source, logger)
				getTendermintHex(
					sourceChain,
					int64(cfg.Chain.Source.Cache.StartHeight),
					cfg.Chain.Source.Tendermint.ChainName,
					logger,
				)
				logger.Info("2. init dest chain")
				getETHJson(&cfg.Chain.Dest, sourceChain, logger)
			}

			if cfg.Chain.Source.ChainType == ETH && cfg.Chain.Dest.ChainType == Tendermint {
				fmt.Println(channelType, "111111")
				logger := log.WithFields(log.Fields{
					"source_chain": &cfg.Chain.Source.Eth.ChainName,
					"dest_chain":   &cfg.Chain.Dest.Tendermint.ChainName,
				})
				logger.Info("1. init dest  chain")
				destChain := tendermintCreateClientFiles(&cfg.Chain.Dest, logger)
				getTendermintHex(
					destChain,
					int64(cfg.Chain.Dest.Cache.StartHeight),
					cfg.Chain.Dest.Tendermint.ChainName,
					logger,
				)
				logger.Info("2. init source chain")
				getETHJson(&cfg.Chain.Source, destChain, logger)
			}

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
	startHeight := latestHeight - 10
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
	hash := common.FromHex(cfg.Eth.Contracts.Packet.Addr)
	clientState := &tibceth.ClientState{
		Header:          header.ToHeader(),
		ChainId:         cfg.Eth.ChainID,
		ContractAddress: hash,
		TrustingPeriod:  60 * 60 * 24 * 7,
		TimeDelay:       0,
		BlockDelay:      7,
	}
	consensusState := &tibceth.ConsensusState{
		Timestamp: header.Time,
		Number:    number,
		Root:      header.Root[:],
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

func getTendermintHex(
	client coresdk.Client,
	height int64,
	chainName string,
	logger *log.Entry) {
	type TrustLevel struct {
		Numerator   int `json:"numerator"`
		Denominator int `json:"denominator"`
	}

	type LatestHeight struct {
		RevisionNumber int   `json:"revisionNumber"`
		RevisionHeight int64 `json:"revisionHeight"`
	}

	type MerklePrefix struct {
		KeyPrefix []byte `json:"keyPrefix"`
	}

	// Tendermint Client State In  ETH
	type TendermintClientState struct {
		ChainID         string       `json:"chainId"`
		TrustLevel      TrustLevel   `json:"trustLevel"`
		TrustingPeriod  int          `json:"trustingPeriod"`
		UnbondingPeriod int          `json:"unbondingPeriod"`
		MaxClockDrift   int          `json:"maxClockDrift"`
		LatestHeight    LatestHeight `json:"latestHeight"`
		MerklePrefix    MerklePrefix `json:"merklePrefix"`
		TimeDelay       int          `json:"timeDelay"`
	}

	type Timestamp struct {
		Secs  int64 `json:"secs"`
		Nanos int64 `json:"nanos"`
	}

	// Tendermint Consensus State In  ETH
	type TendermintConsensusState struct {
		Timestamp          Timestamp `json:"timestamp"`
		Root               string    `json:"root"`
		NextValidatorsHash string    `json:"nextValidatorsHash"`
	}

	blockRes, err := client.QueryBlock(height)
	if err != nil {
		logger.Fatal("QueryBlock fail:  ", err)
	}
	tmHeader := blockRes.Block.Header
	prHeader := tendermint.TmHeaderToPrHeader(tmHeader)

	revisionNumber := int(prHeader.GetHeight().GetRevisionNumber())
	revisionHeight := prHeader.GetHeight().GetRevisionHeight()

	clientState := &TendermintClientState{
		ChainID: tmHeader.ChainID,
		TrustLevel: TrustLevel{
			Numerator:   1,
			Denominator: 3,
		},
		TrustingPeriod:  100 * 24 * 60 * 60,
		UnbondingPeriod: 1814400,
		MaxClockDrift:   10,
		LatestHeight: LatestHeight{
			RevisionNumber: revisionNumber,
			RevisionHeight: int64(revisionHeight),
		},
		MerklePrefix: MerklePrefix{
			KeyPrefix: []byte("tibc"),
		},
		TimeDelay: 0,
	}

	//ConsensusState
	consensusState := TendermintConsensusState{
		Timestamp: Timestamp{
			Secs:  tmHeader.Time.Unix(),
			Nanos: 0,
		},
		Root:               tmHeader.AppHash.String(),
		NextValidatorsHash: tmHeader.NextValidatorsHash.String(),
	}

	clientStateBytes, err := json.Marshal(clientState)
	if err != nil {
		logger.Fatal("marshal eth clientState error: ", err)
	}
	// write file

	clientStateFilename := fmt.Sprintf("%s_clientState.json", chainName)
	writeCreateClientFiles(clientStateFilename, string(clientStateBytes))

	clientStateFilename2 := fmt.Sprintf("%s_clientState.txt", chainName)
	writeCreateClientFiles(clientStateFilename2, hexutil.Encode(clientStateBytes)[2:])
	fmt.Println("clientState: ", hexutil.Encode(clientStateBytes)[2:])

	consensusStateBytes, err := json.Marshal(consensusState)
	if err != nil {
		logger.Fatal(err)
	}

	consensusStateFilename := fmt.Sprintf("%s_consensusState.json", chainName)
	writeCreateClientFiles(consensusStateFilename, string(consensusStateBytes))
	consensusStateFilename2 := fmt.Sprintf("%s_consensusState.txt", chainName)
	writeCreateClientFiles(consensusStateFilename2, hexutil.Encode(consensusStateBytes)[2:])

	fmt.Println("consensusState: ", hexutil.Encode(consensusStateBytes)[2:])
}

func getTendermintJson(
	client coresdk.Client,
	height int64,
	chainName string,
) {

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

	prHeader := tendermint.TmHeaderToPrHeader(tmHeader)

	lastHeight := tibcclient.NewHeight(
		prHeader.GetHeight().GetRevisionNumber(),
		prHeader.GetHeight().GetRevisionHeight())
	var clientState = &tendermint.ClientState{
		ChainId:         tmHeader.ChainID,
		TrustLevel:      fra,
		TrustingPeriod:  time.Hour * 24 * 70 * 2,
		UnbondingPeriod: time.Hour * 24 * 70 * 3,
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
