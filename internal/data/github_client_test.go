package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/model"
)

func TestNewGithubClient(t *testing.T) {
	t.Run("it should create an instance correctly when GITHUB_TOKEN is defined", func(t *testing.T) {
		t.Setenv("GITHUB_TOKEN", "foo")
		client, err := NewGithubClient(context.Background(), slogger.NewDevNullLogger())
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})
	t.Run("it should return an error when GITHUB_TOKEN is not defined", func(t *testing.T) {
		t.Setenv("GITHUB_TOKEN", "")
		client, err := NewGithubClient(context.Background(), slogger.NewDevNullLogger())
		assert.ErrorIs(t, err, ErrNoGithubToken)
		assert.Nil(t, client)
	})
	t.Run("it should implement the model.GithubClient interface", func(t *testing.T) {
		assert.Implements(t, (*model.GithubClient)(nil), &GithubClient{})
	})
}
