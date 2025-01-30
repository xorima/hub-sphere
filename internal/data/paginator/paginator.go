package paginator

import (
	"context"

	"github.com/google/go-github/v68/github"
)

// List is an interface that returns a list of items for the given type
// for example it could be:
//
//	type listFunc struct {
//		client *github.Client
//	}
//
//	func (l *listFunc) List(ctx context.Context, opt *github.ListOptions) ([]*github.Repository, *github.Response, error) {
//		t, r, err := l.client.Apps.ListRepos(ctx, opt)
//		return t.Repositories, r, err
//	}
type List[T any] func(ctx context.Context, opt *github.ListOptions) ([]T, *github.Response, error)

// Process is a function that processes an item for the given type
// this is optional as the user may not want to process the items so
// they can input a skip function that does nothing
// example:
//
//	func Process(ctx context.Context, item *github.Repository) error {
//		fmt.Println(item.GetName())
//		return nil
//	}
type Process[T any] func(ctx context.Context, item T) error

// RateLimit is a function that handles rate limiting
// it returns a bool to indicate if the pagination should continue
// and an error if the users wishes to return more information/errors
// example:
//
//	func RateLimit(ctx context.Context, resp *github.Response) (bool, error) {
//		return true, nil
//	}
type RateLimit func(ctx context.Context, resp *github.Response) (bool, error)

type Opts struct {
	*github.ListOptions
}

func Paginator[T any](ctx context.Context, listFunc List[T], processFunc Process[T], rateLimitFunc RateLimit, Opts *Opts) ([]T, error) {
	var allItems []T

	opts := listOpts(Opts)

	for {
		items, resp, err := listFunc(ctx, opts)
		if err != nil {
			return allItems, err
		}

		allItems = append(allItems, items...)

		for _, item := range items {
			if err = processFunc(ctx, item); err != nil {
				return allItems, err
			}
		}

		// Handle rate limits
		shouldContinue, err := rateLimitFunc(ctx, resp)
		if err != nil {
			return allItems, err
		}

		if !shouldContinue {
			break
		}

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	return allItems, nil
}

func listOpts(opts *Opts) *github.ListOptions {

	if opts == nil || opts.ListOptions == nil {
		return &github.ListOptions{PerPage: 100, Page: 1}
	}
	if opts.PerPage == 0 {
		opts.PerPage = 100
	}

	return opts.ListOptions
}
