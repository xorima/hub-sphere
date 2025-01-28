package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/app"
	"github.com/xorima/hub-sphere/internal/config"
	"github.com/xorima/hub-sphere/internal/data"
	"github.com/xorima/hub-sphere/internal/manager"
	"github.com/xorima/hub-sphere/internal/output"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hub-sphere",
	Short: "A Cli to help manage github at scale focusing in on user actives",
	Long: `This cli application will allow you to mass approved, handle notifications
	and other pull request and user side github activity`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hub-sphere.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getApp(useName string) *app.App {
	log := slogger.NewLogger(slogger.NewLoggerOpts("hub-sphere", useName))
	cfg, err := config.LoadAppConfig(cfgFile)
	cobra.CheckErr(err)
	client, err := data.NewGithubClient(context.Background(), log)
	cobra.CheckErr(err)
	mgr := manager.NewGithubManager(log, client)
	return app.NewApp(log, cfg, mgr, output.NewConsoleOutput())
}
