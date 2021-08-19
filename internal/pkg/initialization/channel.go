package initialization

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services/channels"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/types/cache"
	metricsmodel "github.com/bianjieai/tibc-relayer-go/internal/pkg/types/mertics"
	"github.com/bianjieai/tibc-relayer-go/tools"
	log "github.com/sirupsen/logrus"
)

const TendermintAndTendermint = "tendermint_and_tendermint"

func ChannelMap(cfg *configs.Config, logger *log.Logger) map[string]channels.IChannel {
	relayerMap := map[string]channels.IChannel{}

	for _, channelType := range cfg.App.ChannelTypes {
		switch channelType {
		case TendermintAndTendermint:

			sourceChain := tendermintChain(&cfg.Chain.Source, logger)
			destChain := tendermintChain(&cfg.Chain.Dest, logger)

			// init source chain channe
			sourceChannel := tendermintToTendermint(cfg, sourceChain, destChain, logger)

			// add error_handler mw
			sourceChannel = channels.NewWriterMW(
				sourceChannel, sourceChain.ChainName(), logger,
				tools.DefaultHomePath, tools.DefaultCacheDirName, cfg.Chain.Source.Cache.Filename,
			)

			// add metric mw
			metricsModel := metricsmodel.NewMetric(sourceChain.ChainName())
			sourceChannel = channels.NewMetricMW(sourceChannel, metricsModel)

			// init dest chain channel
			destChannel := tendermintToTendermint(cfg, destChain, sourceChain, logger)

			// add error_handler mw
			destChannel = channels.NewWriterMW(
				destChannel, destChain.ChainName(), logger,
				tools.DefaultHomePath, tools.DefaultCacheDirName, cfg.Chain.Dest.Cache.Filename,
			)

			// add metric mw

			destChannel = channels.NewMetricMW(destChannel, metricsModel)

			relayerMap[sourceChain.ChainName()] = sourceChannel
			relayerMap[destChain.ChainName()] = destChannel

		default:
			logger.WithFields(log.Fields{
				"channel_type": channelType,
			}).Fatal("channel type does not exist")
		}
	}

	return relayerMap
}

func tendermintToTendermint(cfg *configs.Config, sourceChain, destChain repostitory.IChain, logger *log.Logger) channels.IChannel {

	var channel channels.IChannel
	filename := path.Join(tools.DefaultCacheDirName, cfg.Chain.Source.Cache.Filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the file does not exist, the initial height is the startHeight in the configuration

		channel = channels.NewChannel(sourceChain, destChain, cfg.Chain.Source.Cache.StartHeight)
	} else {
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
		channel = channels.NewChannel(sourceChain, destChain, cacheData.LatestHeight)
	}

	return channel
}
