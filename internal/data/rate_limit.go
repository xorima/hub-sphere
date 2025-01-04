package data

import (
	"context"

	"github.com/google/go-github/v68/github"
)

type rateLimitExit struct {
}

func (r *rateLimitExit) RateLimit(ctx context.Context, resp *github.Response) (bool, error) {
	if resp.Rate.Remaining <= 1 {
		return false, nil
	}
	return true, nil
}
