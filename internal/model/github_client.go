package model

import (
	"context"

	"github.com/google/go-github/v71/github"

	"github.com/xorima/hub-sphere/internal/data/paginator"
)

type PullRequest interface {
	GetTitle() string
	GetHTMLURL() string
}

type Repository interface {
	GetName() string
}

type GithubClient interface {
	ListPullRequests(ctx context.Context, owner, repo string, processFunc paginator.Process[*github.PullRequest]) ([]PullRequest, error)
	ListRepositoriesByOrg(ctx context.Context, owner string, processFunc paginator.Process[*github.Repository]) ([]Repository, error)
}
