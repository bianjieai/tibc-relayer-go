package relayer

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/initialization"
)

func Serve(cfg *configs.Config) {

	logger := initialization.Logger(cfg)
	logger.Info("1. service init relayers ")
	channelMap := initialization.ChannelMap(cfg, logger)
	logger.Info("2. service init listener ")
	listener := services.NewListener(channelMap, logger)
	logger.Info("3. service start ")
	logger.Fatal(listener.Listen())
}
