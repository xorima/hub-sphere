package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/app"
	"github.com/xorima/hub-sphere/internal/config"
	"github.com/xorima/hub-sphere/internal/data"
	"github.com/xorima/hub-sphere/internal/manager"
)

// prCmd represents the pr command
var prListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pull requests",
	Long: `Lists pull requests based on a saved filter or a given cli argument
	set`,
	Run: func(cmd *cobra.Command, args []string) {
		log := slogger.NewLogger(slogger.NewLoggerOpts("hub-sphere", "pr"))
		cfg, err := config.LoadAppConfig(cfgFile)
		cobra.CheckErr(err)
		client, err := data.NewGithubClient(context.Background(), log)
		cobra.CheckErr(err)
		mgr := manager.NewGithubManager(log, client)
		a := app.NewApp(log, cfg, mgr)
		a.OpenPullRequests()
	},
}

func init() {
	prCmd.AddCommand(prListCmd)
}
