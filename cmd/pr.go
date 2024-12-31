package cmd

import (
	"github.com/spf13/cobra"
)

// prCmd represents the pr command
var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Pull Request Management",
	Long: `All actions related to pull requests sit under this command.
	see sub commands for detail`,
}

func init() {
	rootCmd.AddCommand(prCmd)
}
