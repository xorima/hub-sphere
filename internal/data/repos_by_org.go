package data

import (
	"context"

	"github.com/google/go-github/v71/github"

	"github.com/xorima/hub-sphere/internal/data/paginator"
	"github.com/xorima/hub-sphere/internal/model"
)

func listOrgs(client *github.Client, owner string) paginator.List[*github.Repository] {
	return func(ctx context.Context, opt *github.ListOptions) ([]*github.Repository, *github.Response, error) {
		return client.Repositories.ListByOrg(ctx, owner, &github.RepositoryListByOrgOptions{Type: "all", ListOptions: *opt})
	}
}

func (c *GithubClient) ListRepositoriesByOrg(ctx context.Context, owner string, processFunc paginator.Process[*github.Repository]) ([]model.Repository, error) {
	items, err := paginator.Paginator[*github.Repository](ctx, listOrgs(c.client, owner), processFunc, rateLimitExit, &paginator.Opts{
		ListOptions: &github.ListOptions{PerPage: 50, Page: 1},
	})
	if err != nil {
		return nil, err
	}
	var result []model.Repository
	for _, item := range items {
		result = append(result, item)
	}
	return result, nil
}
