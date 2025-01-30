package model

import (
	"context"

	"github.com/xorima/hub-sphere/internal/config"
)

type RepositoryPR struct {
	RepoName string
	Pr       []PullRequest
}

type GithubManager interface {
	OpenPullRequests(ctx context.Context, filter config.Filter) ([]RepositoryPR, error)
}
