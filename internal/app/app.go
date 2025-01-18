package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/xorima/hub-sphere/internal/config"
	"github.com/xorima/hub-sphere/internal/manager"
)

type App struct {
	config *config.HubSphere
	log    *slog.Logger
	mgr    *manager.GithubManager
}

func NewApp(log *slog.Logger, cfg *config.HubSphere, mgr *manager.GithubManager) *App {
	return &App{log: log, config: cfg, mgr: mgr}
}

func (a *App) OpenPullRequests() {
	// make this more efficient by having mgr pass in the processor so we can live process each object
	// and have app pass in a formatter so it can own how they are outputted keeping separation of concerns.
	// once this is done we need to add in filtering
	// and tests, unit test everything... + interfaces to make testing possible...
	repos, err := a.mgr.OpenPullRequests(context.Background(), "sous-chefs")
	if err != nil {
		a.log.Error("error occurred, exiting")
		return
	}
	for _, repo := range repos {
		if len(repo.Pr) == 0 {
			continue
		}
		fmt.Printf("%s: PRs\n", repo.RepoName)
		for _, pr := range repo.Pr {
			fmt.Printf("\t - %s : %s\n", pr.GetTitle(), pr.GetHTMLURL())
		}
	}
}
