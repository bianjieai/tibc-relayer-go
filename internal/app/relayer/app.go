package relayer

import (
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services/channels"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/initialization"
)

func Serve(cfg *configs.Config) {

	logger := initialization.Logger(cfg)
	logger.Info("1. service init relayers ")
	channelMap := initialization.ChannelMap(cfg, logger)
	logger.Info("2. service init listener & TaskPerformer")
	listener := services.NewListener(channelMap, logger)
	logger.Info("3. service start & crontab start")
	go runTask(channelMap, logger)
	go runMetricHandler(cfg, logger)
	logger.Fatal(listener.Listen())
}

func runTask(channelMap map[string]channels.IChannel, logger *log.Logger) {
	for channelName := range channelMap {
		// execute every x hours
		err := gocron.Every(channelMap[channelName].UpdateClientFrequency()).Hours().Do(func() {
			err := channelMap[channelName].UpdateClient()
			logger.Error(err)
		})
		logger.Error(err)
	}
	_, nextTime := gocron.NextRun()
	logger.WithFields(log.Fields{"next_time": nextTime}).Info()
	<-gocron.Start()
}

func runMetricHandler(cfg *configs.Config, logger *log.Logger) {
	logger.Info("scanner metric start: addr ", cfg.App.MetricAddr)
	metricMux := http.NewServeMux()
	metricMux.Handle("/metrics", promhttp.Handler())
	logger.Fatal(http.ListenAndServe(cfg.App.MetricAddr, metricMux))
}
