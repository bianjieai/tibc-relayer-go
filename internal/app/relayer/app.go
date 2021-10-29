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
	s := gocron.NewScheduler()
	for channelName := range channelMap {
		// execute every x hours
		//err := gocron.Every(channelMap[channelName].UpdateClientFrequency()).Hours().Do(func()
		channel := channelMap[channelName]
		doFunc := func(channel channels.IChannel) {
			err := channel.UpdateClient()
			if err != nil {
				logger.Error(err)
			}
		}
		updateClientFrequency := channel.UpdateClientFrequency()
		s.Every(updateClientFrequency).Seconds().Do(doFunc, channel)
	}

	_, nextTime := s.NextRun()
	logger.WithFields(log.Fields{"next_time": nextTime}).Info()
	<-s.Start()
}

func runMetricHandler(cfg *configs.Config, logger *log.Logger) {
	logger.Info("scanner metric start: addr ", cfg.App.MetricAddr)
	metricMux := http.NewServeMux()
	metricMux.Handle("/metrics", promhttp.Handler())
	logger.Fatal(http.ListenAndServe(cfg.App.MetricAddr, metricMux))
}
