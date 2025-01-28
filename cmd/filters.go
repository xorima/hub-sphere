package cmd

import (
	"github.com/spf13/cobra"
)

// filtersCmd represents the filters command
var filtersCmd = &cobra.Command{
	Use:   "filters",
	Short: "Holds all commands related to config filters",
}

func init() {
	configCmd.AddCommand(filtersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filtersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filtersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
