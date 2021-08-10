package tools

import (
	"os"
	"path/filepath"
)

const CfgDirName = ".tibc-relayer"

var (
	UserDir, _       = os.UserHomeDir()
	DefaultConfigDir = filepath.Join(UserDir, CfgDirName)
)
