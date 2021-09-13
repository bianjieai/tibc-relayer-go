package relayer

import (
	"github.com/jasonlvhit/gocron"
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
	logger.Fatal(listener.Listen())
}

func runTask(channelMap map[string]channels.IChannel, logger *log.Logger) {
	for channelName := range channelMap {
		// execute every x hours
		gocron.Every(channelMap[channelName].UpdateClientFrequency()).Hours().Do(func() {
			channelMap[channelName].UpdateClient()
		})
	}
	_, nextTime := gocron.NextRun()
	logger.WithFields(log.Fields{"next_time": nextTime}).Info()
	<-gocron.Start()
}
