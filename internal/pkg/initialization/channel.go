package initialization

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services/relayer"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	log "github.com/sirupsen/logrus"
)

const WenchangToBsnHub = "wenchang_to_bsnhub"
const BsnHubToWenchang = "bsnhub_to_wenchang"

func Channels(cfg *configs.Config, logger *log.Logger) map[string]relayer.IRelayer {
	relayerMap := map[string]relayer.IRelayer{}

	for _, sourceToDest := range cfg.App.SourceToDest {
		switch sourceToDest {
		case WenchangToBsnHub:
			relayerMap[sourceToDest] = WenchangToBsnRelayer(cfg, logger)
		case BsnHubToWenchang:
			relayerMap[sourceToDest] = BsnHubToWenchangRelayer(cfg, logger)
		default:
			logger.WithFields(log.Fields{
				"channel": cfg.App.SourceToDest,
			}).Fatal("channel does not exist")
		}
	}

	return relayerMap
}
