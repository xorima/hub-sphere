package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pull requests",
	Long: `Lists pull requests based on a saved filter or a given cli argument
	set`,
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(getApp("pr-list").OpenPullRequests(context.Background()))
	},
}

func init() {
	prCmd.AddCommand(prListCmd)
}
