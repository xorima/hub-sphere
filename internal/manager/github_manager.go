package manager

import (
	"context"
	"log/slog"

	"github.com/google/go-github/v71/github"
	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/config"
	"github.com/xorima/hub-sphere/internal/data"
	"github.com/xorima/hub-sphere/internal/model"
	"github.com/xorima/hub-sphere/internal/output"
)

type GithubManager struct {
	log       *slog.Logger
	client    model.GithubClient
	outputter model.Outputter
}

func NewGithubManager(log *slog.Logger, client model.GithubClient) *GithubManager {
	return &GithubManager{log: log, client: client, outputter: output.NewConsoleOutput()}
}

func (m *GithubManager) OpenPullRequests(ctx context.Context, filter config.Filter) ([]model.RepositoryPR, error) {
	repositories, err := m.client.ListRepositoriesByOrg(ctx, filter.Owner, data.ProcessDoNothing[*github.Repository]())
	if err != nil {
		m.log.Error("get by org error", slogger.ErrorAttr(err), slog.String("owner", filter.Owner))
		return nil, err
	}
	var resp []model.RepositoryPR
	for _, r := range repositories {
		var tmp = model.RepositoryPR{
			RepoName: r.GetName(),
		}
		prs, err := m.client.ListPullRequests(ctx, filter.Owner, r.GetName(), data.ProcessDoNothing[*github.PullRequest]())
		if err != nil {
			m.log.Error("list prs error", slogger.ErrorAttr(err), slog.String("owner", filter.Owner), slog.String("repo", r.GetName()))
			return resp, err
		}
		tmp.Pr = append(tmp.Pr, prs...)
		resp = append(resp, tmp)
	}
	return resp, nil
}
