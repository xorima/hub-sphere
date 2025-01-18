package manager

import (
	"context"
	"testing"

	"github.com/google/go-github/v68/github"
	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/model"
)

const (
	ownerForTest = "sous-chefs"
	repoForTest  = "apache2"
)

type mockGithubClient struct {
	repositories    []*model.Repository
	repositoriesErr error
	pullRequests    []*model.PullRequest
	pullRequestsErr error
}

func newMockGithubClient() *mockGithubClient {
	return &mockGithubClient{
		repositories: make([]*model.Repository, 0),
		pullRequests: make([]*model.PullRequest, 0),
	}
}

func (m *mockGithubClient) ListRepositoriesByOrg(ctx context.Context, owner string) ([]*model.Repository, error) {
	return m.repositories, m.repositoriesErr
}

func (m *mockGithubClient) ListPullRequests(ctx context.Context, owner, repo string) ([]*model.PullRequest, error) {
	return m.pullRequests, m.pullRequestsErr
}

func TestNewGithubManager(t *testing.T) {
	t.Run("it should create an instance of the github manager", func(t *testing.T) {
		mgr := NewGithubManager(slogger.NewDevNullLogger(), newMockGithubClient())
		assert.NotNil(t, mgr)
	})
}

func TestGithubManagerOpenPullRequests(t *testing.T) {
	t.Run("it should return an error if it cannot list the repositories by org", func(t *testing.T) {
		c := newMockGithubClient()
		c.repositoriesErr = assert.AnError
		mgr := NewGithubManager(slogger.NewDevNullLogger(), c)
		res, err := mgr.OpenPullRequests(context.Background(), ownerForTest)
		assert.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, res)
	})
	t.Run("it should return nothing if no repositories are found", func(t *testing.T) {
		c := newMockGithubClient()
		mgr := NewGithubManager(slogger.NewDevNullLogger(), c)
		res, err := mgr.OpenPullRequests(context.Background(), ownerForTest)
		assert.NoError(t, err)
		assert.Nil(t, res)
	})
	t.Run("it should return an error if it cannot list the pull requests", func(t *testing.T) {
		c := newMockGithubClient()
		c.repositories = append(c.repositories, &model.Repository{Name: github.Ptr[string](repoForTest)})
		c.pullRequestsErr = assert.AnError
		mgr := NewGithubManager(slogger.NewDevNullLogger(), c)
		res, err := mgr.OpenPullRequests(context.Background(), ownerForTest)
		assert.ErrorIs(t, err, assert.AnError)
		assert.Nil(t, res)
	})
	t.Run("it should return the open pull requests that belong to the repositories", func(t *testing.T) {
		c := newMockGithubClient()
		c.repositories = append(c.repositories, &model.Repository{Name: github.Ptr[string](repoForTest)})
		c.pullRequests = append(c.pullRequests, &model.PullRequest{ID: github.Ptr[int64](1)})
		mgr := NewGithubManager(slogger.NewDevNullLogger(), c)
		res, err := mgr.OpenPullRequests(context.Background(), ownerForTest)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res, 1)
		assert.Equal(t, repoForTest, res[0].RepoName)
		assert.Len(t, res[0].Pr, 1)
		assert.Equal(t, int64(1), res[0].Pr[0].GetID())
	})
}
