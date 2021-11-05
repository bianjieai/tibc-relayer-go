package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/bianjieai/tibc-relayer-go/cmd/handlers"

	"github.com/spf13/cobra"
)

var (
	endHeight uint64
	batchCmd  = &cobra.Command{
		Use:   "batch",
		Short: "Manual batch send update eth tx",
		Run: func(cmd *cobra.Command, args []string) {
			batchUpdateETHClient()
		},
	}
)

func init() {
	rootCmd.AddCommand(batchCmd)
	batchCmd.Flags().StringVarP(&localConfig, "CONFIG", "c", "", "config path: /opt/local.toml")
	batchCmd.Flags().Uint64VarP(&endHeight, "END", "e", 0, "ethereum ending height")
}

func batchUpdateETHClient() {
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

	handlers.BatchUpdateETHClient(config, endHeight)
}
