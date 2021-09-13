package initialization

import (
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
)

func Logger(cfg *configs.Config) *log.Logger {
	logger := log.New()
	if cfg.App.Env == "prod" {
		logger.SetFormatter(&log.JSONFormatter{})
	}
	switch cfg.App.LogLevel {
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
	return logger
}
