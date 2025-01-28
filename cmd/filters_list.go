package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available filters from the config file",
	Long: `This will show all available filters and their configuration
	which is useful for when you wish to know what you can filter by from
	your config easily`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(getApp("config-filter-list").AvailableFilters())
	},
}

func init() {
	filtersCmd.AddCommand(listCmd)
}
