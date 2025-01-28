package data

import (
	"context"

	"github.com/google/go-github/v68/github"
	"github.com/hsbc/go-api-pagination/pagination"

	"github.com/xorima/hub-sphere/internal/model"
)

type listReposByOrg struct {
	owner  string
	client *github.Client
}

func (l *listReposByOrg) List(ctx context.Context, opt *github.ListOptions) ([]*github.Repository, *github.Response, error) {
	t, r, err := l.client.Repositories.ListByOrg(ctx, l.owner, &github.RepositoryListByOrgOptions{Type: "all", ListOptions: *opt})
	return t, r, err
}

func (l *listReposByOrg) Process(ctx context.Context, item *github.Repository) error {
	return nil
}

func (c *GithubClient) ListRepositoriesByOrg(ctx context.Context, owner string) ([]model.Repository, error) {
	o := &listReposByOrg{owner: owner, client: c.client}
	items, err := pagination.Paginator[*github.Repository](ctx, o, o, &rateLimitExit{}, &pagination.PaginatorOpts{
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
