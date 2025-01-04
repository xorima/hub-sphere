package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestGithubClientListRepositoriesByOrg(t *testing.T) {
	t.Run("it should return all the repositories from a paginated list", func(t *testing.T) {
		c, r, err := newVcrGithubClient(t, "fixtures/repository-list", AccessToken)
		assert.NoError(t, err)
		defer func(r *recorder.Recorder) {
			assert.NoError(t, r.Stop())
		}(r)
		repos, err := c.ListRepositoriesByOrg(context.Background(), "sous-chefs")
		assert.NoError(t, err)
		assert.Len(t, repos, 184)
	})
	t.Run("it should return an error if the upstream errors", func(t *testing.T) {
		c, err := NewGithubClient(context.Background(), slogger.NewDevNullLogger(), WithTransport(&mockTransport{}))
		assert.NoError(t, err)
		repos, err := c.ListRepositoriesByOrg(context.Background(), "sous-chefs")
		assert.ErrorIs(t, err, assert.AnError)
		assert.Len(t, repos, 0)
	})
}
