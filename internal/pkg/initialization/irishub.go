package initialization

import (
	"os"
	"path"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services/relayer"
	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/tools"
)

func IrisHubRelayer(cfg *configs.Config) relayer.IRelayer {
	cacheDir := path.Join(tools.DefaultConfigDir, cfg.Chain.IrisHub.Cache.Dir)
	filename := path.Join(cacheDir, cfg.Chain.IrisHub.Cache.Filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// todo
		// If the file does not exist, the initial height is the startHeight in the configuration
		return nil
	} else {
		// todo
		// If the file exists, the initial height is the latest_height in the file
	}
	return nil
}
