package relayer

import (
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/initialization"
)

func Serve(cfg *configs.Config) {

	logger := initialization.Logger(cfg)
	logger.Info("1. service init ")

}
