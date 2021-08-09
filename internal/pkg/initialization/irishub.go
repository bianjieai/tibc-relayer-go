package initialization

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/services/relayer"
	"os"
	"path"
	"path/filepath"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
)

var (
	userDir, _ = os.UserHomeDir()
	homeDir    = filepath.Join(userDir, ".tibc-relayer")
)

func IrisHubRelayer(cfg *configs.Config) relayer.IRelayer {
	cacheDir := path.Join(homeDir, cfg.Chain.IrisHub.Cache.Dir)
	filename := path.Join(cacheDir, cfg.Chain.IrisHub.Cache.Filename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// todo
		return nil
	} else {
		//todo
	}
	return nil
}
