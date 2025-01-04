package data

import (
	"context"

	"github.com/google/go-github/v68/github"
	"github.com/hsbc/go-api-pagination/pagination"
)

type PullRequest = github.PullRequest

type listPullRequests struct {
	owner  string
	repo   string
	client *github.Client
}

func (l *listPullRequests) List(ctx context.Context, opt *github.ListOptions) ([]*github.PullRequest, *github.Response, error) {
	t, r, err := l.client.PullRequests.List(ctx, l.owner, l.repo, &github.PullRequestListOptions{State: "open", ListOptions: *opt})
	return t, r, err
}

func (l *listPullRequests) Process(ctx context.Context, item *github.PullRequest) error {
	return nil
}

func (c *GithubClient) ListPullRequests(ctx context.Context, owner, repo string) ([]*PullRequest, error) {
	o := &listPullRequests{owner: owner, repo: repo, client: c.client}
	items, err := pagination.Paginator[*github.PullRequest](ctx, o, o, &rateLimitExit{}, &pagination.PaginatorOpts{
		ListOptions: &github.ListOptions{PerPage: 50, Page: 1},
	})
	return items, err
}
