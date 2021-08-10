package tools

import (
	"os"
	"path/filepath"
)

const DefaultHomeDirName = ".tibc-relayer"
const DefaultConfigDirName = "configs"
const DefaultConfigName = "config.toml"
const DefaultCacheDirName = "cache"

var (
	UserDir, _      = os.UserHomeDir()
	DefaultHomePath = filepath.Join(UserDir, DefaultHomeDirName)
)
