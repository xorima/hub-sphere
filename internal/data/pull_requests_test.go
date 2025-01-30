package data

import (
	"context"
	"testing"

	"github.com/google/go-github/v68/github"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestListPullRequests(t *testing.T) {
	t.Run("it should return the open pull requests", func(t *testing.T) {
		c, r, err := newVcrGithubClient(t, "fixtures/pull-request-list", AccessToken)
		assert.NoError(t, err)
		defer func(r *recorder.Recorder) {
			assert.NoError(t, r.Stop())
		}(r)
		repos, err := c.ListPullRequests(context.Background(), "sous-chefs", "java", ProcessDoNothing[*github.PullRequest]())
		assert.NoError(t, err)
		assert.Len(t, repos, 1)
	})
	t.Run("it should return an error if the upstream errors", func(t *testing.T) {
		c, err := NewGithubClient(context.Background(), slogger.NewDevNullLogger(), WithTransport(&mockTransport{}))
		assert.NoError(t, err)
		repos, err := c.ListPullRequests(context.Background(), "sous-chefs", "java", ProcessDoNothing[*github.PullRequest]())
		assert.ErrorIs(t, err, assert.AnError)
		assert.Len(t, repos, 0)

	})
}
