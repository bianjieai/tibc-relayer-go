package handlers

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	tibcclient "github.com/bianjieai/tibc-sdk-go/client"
	"github.com/bianjieai/tibc-sdk-go/commitment"
	"github.com/bianjieai/tibc-sdk-go/tendermint"
	coresdk "github.com/irisnet/core-sdk-go"
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
	tmtypes "github.com/tendermint/tendermint/types"
)

const TendermintAndTendermint = "tendermint_and_tendermint"

const tibcTendermintMerklePrefix = "tibc"
const tibcTendermintRoot = "app_hash"

const clientStatePrefix = `{"@type":"/tibc.lightclients.tendermint.v1.ClientState",`
const consensusStatePrefix = `{"@type":"/tibc.lightclients.tendermint.v1.ConsensusState",`

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
		}
	}
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
	//及时关闭file句柄
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
