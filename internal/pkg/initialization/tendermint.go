package initialization

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services/relayer"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/cache"
	"github.com/bianjieai/tibc-relayer-go/tools"
	log "github.com/sirupsen/logrus"
)

func WenchangToBsnRelayer(cfg *configs.Config, logger *log.Logger) relayer.IRelayer {

	sourceChain := wenchangChain(cfg, logger)
	destChain := bsnHubChain(cfg, logger)

	filename := path.Join(tools.DefaultCacheDirName, cfg.Chain.Wenchang.Cache.Filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the file does not exist, the initial height is the startHeight in the configuration

		return relayer.NewRelayer(sourceChain, destChain, cfg.Chain.Wenchang.Cache.StartHeight)
	}

	// If the file exists, the initial height is the latest_height in the file

	file, err := os.Open(filename)
	if err != nil {
		logger.Fatal("read cache file err: ", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Fatal("read cache file err: ", err)
	}

	cacheData := &cache.Data{}
	err = json.Unmarshal(content, cacheData)
	if err != nil {
		logger.Fatal("read cache file unmarshal err: ", err)
	}

	return relayer.NewRelayer(sourceChain, destChain, cacheData.LatestHeight)
}

func BsnHubToWenchangRelayer(cfg *configs.Config, logger *log.Logger) relayer.IRelayer {

	sourceChain := wenchangChain(cfg, logger)
	destChain := bsnHubChain(cfg, logger)

	filename := path.Join(tools.DefaultCacheDirName, cfg.Chain.Wenchang.Cache.Filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the file does not exist, the initial height is the startHeight in the configuration

		return relayer.NewRelayer(sourceChain, destChain, cfg.Chain.Wenchang.Cache.StartHeight)
	}

	// If the file exists, the initial height is the latest_height in the file

	file, err := os.Open(filename)
	if err != nil {
		logger.Fatal("read cache file err: ", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Fatal("read cache file err: ", err)
	}

	cacheData := &cache.Data{}
	err = json.Unmarshal(content, cacheData)
	if err != nil {
		logger.Fatal("read cache file unmarshal err: ", err)
	}

	return relayer.NewRelayer(sourceChain, destChain, cacheData.LatestHeight)
}

func wenchangChain(cfg *configs.Config, logger *log.Logger) repostitory.IChain {

	logger.WithFields(log.Fields{
		"chain_name": cfg.Chain.Wenchang.ChainName,
	}).Info(" init chain start")

	chainCfg := repostitory.NewTerndermintConfig()
	chainCfg.ChainID = cfg.Chain.Wenchang.ChainID
	chainCfg.GrpcAddr = cfg.Chain.Wenchang.GrpcAddr
	chainCfg.RPCAddr = cfg.Chain.Wenchang.RPCAddr
	chainRepo, err := repostitory.NewTendermintClient(cfg.Chain.Wenchang.ChainName, chainCfg)
	if err != nil {
		logger.WithFields(log.Fields{
			"chain_name": cfg.Chain.Wenchang.ChainName,
		}).Fatal("failed to init chain")
	}

	return chainRepo
}

func bsnHubChain(cfg *configs.Config, logger *log.Logger) repostitory.IChain {
	logger.WithFields(log.Fields{
		"chain_name": cfg.Chain.BsnHub.ChainName,
	}).Info(" init chain start")

	chainCfg := repostitory.NewTerndermintConfig()
	chainCfg.ChainID = cfg.Chain.BsnHub.ChainID
	chainCfg.GrpcAddr = cfg.Chain.BsnHub.GrpcAddr
	chainCfg.RPCAddr = cfg.Chain.BsnHub.RPCAddr
	chainRepo, err := repostitory.NewTendermintClient(cfg.Chain.BsnHub.ChainName, chainCfg)
	if err != nil {
		logger.WithFields(log.Fields{
			"chain_name": cfg.Chain.BsnHub.ChainName,
		}).Fatal("failed to init chain")
	}

	return chainRepo
}
