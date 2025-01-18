package data

import (
	"context"
	"testing"

	"github.com/google/go-github/v68/github"
	"github.com/stretchr/testify/assert"
)

// This isn't the right way to test this, but it is easy for a function we will not amend at this
// moment in time
func TestRateLimitExitRateLimit(t *testing.T) {
	t.Run("it should return true when there is rate limit left", func(t *testing.T) {
		r := &rateLimitExit{}
		limit, err := r.RateLimit(context.Background(), &github.Response{Rate: github.Rate{Remaining: 10}})
		assert.NoError(t, err)
		assert.True(t, limit)
	})
	t.Run("it should return false when there is no rate limit remaining", func(t *testing.T) {
		r := &rateLimitExit{}
		limit, err := r.RateLimit(context.Background(), &github.Response{Rate: github.Rate{Remaining: 0}})
		assert.NoError(t, err)
		assert.False(t, limit)
	})
}
