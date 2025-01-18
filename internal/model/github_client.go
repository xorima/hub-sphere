package model

import (
	"context"

	"github.com/google/go-github/v68/github"
)

type Repository = github.Repository
type PullRequest = github.PullRequest

type GithubClient interface {
	ListPullRequests(ctx context.Context, owner, repo string) ([]*PullRequest, error)
	ListRepositoriesByOrg(ctx context.Context, owner string) ([]*Repository, error)
}
