package manager

import (
	"context"
	"log/slog"

	"github.com/xorima/slogger"

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

func (m *GithubManager) OpenPullRequests(ctx context.Context, owner string) ([]model.RepositoryPR, error) {
	repositories, err := m.client.ListRepositoriesByOrg(ctx, owner)
	if err != nil {
		m.log.Error("get by org error", slogger.ErrorAttr(err), slog.String("owner", owner))
		return nil, err
	}
	var resp []model.RepositoryPR
	for _, r := range repositories {
		var tmp = model.RepositoryPR{
			RepoName: r.GetName(),
		}
		prs, err := m.client.ListPullRequests(ctx, owner, r.GetName())
		if err != nil {
			m.log.Error("list prs error", slogger.ErrorAttr(err), slog.String("owner", owner), slog.String("repo", r.GetName()))
			return resp, err
		}

		tmp.Pr = append(tmp.Pr, prs...)
		resp = append(resp, tmp)
	}
	return resp, nil
}
