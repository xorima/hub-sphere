package model

import (
	"context"
)

type PullRequest interface {
	GetTitle() string
	GetHTMLURL() string
}

type Repository interface {
	GetName() string
}

type GithubClient interface {
	ListPullRequests(ctx context.Context, owner, repo string) ([]PullRequest, error)
	ListRepositoriesByOrg(ctx context.Context, owner string) ([]Repository, error)
}
