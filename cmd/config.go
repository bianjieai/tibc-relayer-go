package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bianjieai/tibc-relayer-go/cmd/handlers"
	"github.com/spf13/cobra"
)

var (
	userDir, _       = os.UserHomeDir()
	defaultConfigDir = filepath.Join(userDir, ".tibc-relayer")
)

var (
	home string

	configCmd = &cobra.Command{
		Use:     "config",
		Aliases: []string{"cfg"},
		Short:   "manage configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			createConfig()
		},
	}
	configInitCmd = &cobra.Command{
		Use:   "init",
		Short: "init configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			initConfig()
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configInitCmd)
	configCmd.Flags().StringVarP(&home, "path", "p", "", "config path: .relayer")
}

func createConfig() {
	if home == "" {
		fmt.Println("please enter dir, for example: relayer cfg -p .relayer ")
		return
	}
	handlers.ConfigInit(home)
}

func initConfig() {
	handlers.ConfigInit(defaultConfigDir)
}
