package handlers

import (
	"os"
	"path"

	"github.com/bianjieai/tibc-relayer-go/internal/pkg/configs"
	"github.com/bianjieai/tibc-relayer-go/tools"
	"github.com/pelletier/go-toml"
)

func ConfigInit(home string) error {
	cfgDir := path.Join(home, tools.DefaultConfigDirName)
	cfgPath := path.Join(cfgDir, tools.DefaultConfigName)
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		// And the home folder doesn't exist
		if _, err := os.Stat(home); os.IsNotExist(err) {
			// Create the home folder
			if err = os.Mkdir(home, os.ModePerm); err != nil {
				return err
			}
		}
		// Create the home config folder
		if err = os.Mkdir(cfgDir, os.ModePerm); err != nil {
			return err
		}
	}

	// Then create the file...
	f, err := os.Create(cfgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// And write the default config to that location...
	if _, err = f.Write(defaultConfig()); err != nil {
		return err
	}

	// And return no error...
	return nil
}

func defaultConfig() []byte {
	cfg := configs.NewConfig()
	cfg.App.MetricAddr = "0.0.0.0:8083"
	cfg.App.Env = "dev"
	cfg.App.LogLevel = "debug"
	data, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	return data
}
