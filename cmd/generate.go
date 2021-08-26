package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/bianjieai/tibc-relayer-go/cmd/handlers"

	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate the files needed for create client: clientStatus & consensusState",
		Run: func(cmd *cobra.Command, args []string) {
			createClientFiles()
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&localConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
}

func createClientFiles() {
	data, err := ioutil.ReadFile(localConfig)
	if err != nil {
		fmt.Println("Error: get config data error: ", err)
		return
	}
	config, err := readConfig(data)
	if err != nil {
		fmt.Println("Error: read config error: ", err)
		return
	}
	handlers.CreateClientFiles(config)
}
