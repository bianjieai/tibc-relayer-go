package initialization

import (
	coretypes "github.com/irisnet/core-sdk-go/types"
	corestore "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/constant"
)

func tendermintChain(cfg *configs.ChainCfg, logger *log.Logger) repostitory.IChain {

	logger.WithFields(log.Fields{
		"chain_name": cfg.Tendermint.ChainName,
	}).Info(" init chain start")

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
	options := []coretypes.Option{
		coretypes.KeyDAOOption(corestore.NewMemory(corestore.NewMemory(nil))),
		coretypes.TimeoutOption(cfg.Tendermint.RequestTimeout),
		coretypes.ModeOption(coretypes.Commit),
		coretypes.GasOption(cfg.Tendermint.Gas),
		coretypes.CachedOption(true),
	}
	if cfg.Tendermint.Algo != "" {
		options = append(options, coretypes.AlgoOption(cfg.Tendermint.Algo))
	}
	chainCfg.Options = options
	chainRepo, err := repostitory.NewTendermintClient(
		constant.Tendermint,
		cfg.Tendermint.ChainName,
		cfg.Tendermint.UpdateClientFrequency,
		cfg.Tendermint.WhitelistSender,
		chainCfg,
	)
	if err != nil {
		logger.WithFields(log.Fields{
			"chain_name": cfg.Tendermint.ChainName,
			"err_msg":    err,
		}).Fatal("failed to init chain")
	}

	return chainRepo
}
