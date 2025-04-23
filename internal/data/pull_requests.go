package data

import (
	"context"

	"github.com/google/go-github/v71/github"

	"github.com/xorima/hub-sphere/internal/data/paginator"
	"github.com/xorima/hub-sphere/internal/model"
)

func listPullRequests(client *github.Client, owner, repo string) paginator.List[*github.PullRequest] {
	return func(ctx context.Context, opt *github.ListOptions) ([]*github.PullRequest, *github.Response, error) {
		return client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{State: "open", ListOptions: *opt})
	}
}

func (c *GithubClient) ListPullRequests(ctx context.Context, owner, repo string, processFunc paginator.Process[*github.PullRequest]) ([]model.PullRequest, error) {
	items, err := paginator.Paginator[*github.PullRequest](ctx, listPullRequests(c.client, owner, repo), processFunc, rateLimitExit, &paginator.Opts{
		ListOptions: &github.ListOptions{PerPage: 50, Page: 1},
	})
	if err != nil {
		return nil, err
	}
	var result []model.PullRequest
	for _, item := range items {
		result = append(result, item)
	}
	return result, nil
}
