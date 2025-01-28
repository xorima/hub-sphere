package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/config"
	"github.com/xorima/hub-sphere/internal/model"
)

type App struct {
	config *config.HubSphere
	log    *slog.Logger
	mgr    model.GithubManager
	output model.Outputter
}

func NewApp(log *slog.Logger, cfg *config.HubSphere, mgr model.GithubManager, outputter model.Outputter) *App {
	return &App{log: log, config: cfg, mgr: mgr, output: outputter}
}

func (a *App) OpenPullRequests(ctx context.Context) error {
	repos, err := a.mgr.OpenPullRequests(ctx, "sous-chefs")
	if err != nil {
		a.log.ErrorContext(ctx, "error occurred, exiting", slogger.ErrorAttr(err))
		return err
	}
	var entries = make(model.Entries)
	for _, repo := range repos {
		if len(repo.Pr) == 0 {
			continue
		}
		repoName := fmt.Sprintf("%s PRs", repo.RepoName)
		for _, pr := range repo.Pr {
			entries[repoName] = append(entries[repoName], fmt.Sprintf("- %s : %s", pr.GetTitle(), pr.GetHTMLURL()))
		}
	}
	if len(entries) > 0 {
		err = a.output.Write(entries)
		if err != nil {
			a.log.ErrorContext(ctx, "failed to write output", slogger.ErrorAttr(err))
			return err
		}
	}
	return nil
}
