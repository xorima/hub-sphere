package model

import (
	"context"
)

type RepositoryPR struct {
	RepoName string
	Pr       []PullRequest
}

type GithubManager interface {
	OpenPullRequests(ctx context.Context, owner string) ([]RepositoryPR, error)
}
