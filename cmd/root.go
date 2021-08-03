package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "v0.1.0"
	Date    = "2021-07-26"
)

var rootCmd = &cobra.Command{
	Use:   "UCI-Server",
	Short: "UCI-Server for OTCIBR",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
