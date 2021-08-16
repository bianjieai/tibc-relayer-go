package initialization

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	log "github.com/sirupsen/logrus"
)

func tendermintChain(cfg *configs.ChainCfg, logger *log.Logger) repostitory.IChain {

	logger.WithFields(log.Fields{
		"chain_name": cfg.Tendermint.ChainName,
	}).Info(" init chain start")

	chainCfg := repostitory.NewTerndermintConfig()
	chainCfg.ChainID = cfg.Tendermint.ChainID
	chainCfg.GrpcAddr = cfg.Tendermint.GrpcAddr
	chainCfg.RPCAddr = cfg.Tendermint.RPCAddr
	chainRepo, err := repostitory.NewTendermintClient(cfg.Tendermint.ChainName, chainCfg)
	if err != nil {
		logger.WithFields(log.Fields{
			"chain_name": cfg.Tendermint.ChainName,
		}).Fatal("failed to init chain")
	}

	return chainRepo
}
