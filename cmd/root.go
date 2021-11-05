package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "v0.2.0"
	Date    = "2021-11-01"
)

var rootCmd = &cobra.Command{
	Use:   "relayer",
	Short: "relayer for tibc",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
