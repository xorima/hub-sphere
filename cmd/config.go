package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "holds all commands related to config",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
