package app

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"

	"github.com/xorima/hub-sphere/internal/config"
	"github.com/xorima/hub-sphere/internal/model"
	"github.com/xorima/hub-sphere/internal/model/modelmocks"
)

var (
	mockPR   = modelmocks.NewMockPullRequest("chore(deps): update x", "https://github.com/")
	repoName = "example-repo"
)

type mockManager struct {
	pr  []model.RepositoryPR
	err error
}

func (m mockManager) OpenPullRequests(ctx context.Context, filter config.Filter) ([]model.RepositoryPR, error) {
	return m.pr, m.err
}

type mockOutputter struct {
	err     error
	entries model.Entries
}

func (m *mockOutputter) Write(entries model.Entries) error {
	m.entries = entries
	return m.err
}

func TestNewApp(t *testing.T) {
	t.Run("it should return an instance of  the new app", func(t *testing.T) {
		a := NewApp(slogger.NewDevNullLogger(), &config.HubSphere{}, &mockManager{}, &mockOutputter{})
		assert.NotNil(t, a)
	})
}

func TestAppOpenPullRequests(t *testing.T) {
	t.Run("it should return an error if the github manager cannot find open pull requests", func(t *testing.T) {
		a := NewApp(slogger.NewDevNullLogger(), &config.HubSphere{}, &mockManager{err: assert.AnError}, &mockOutputter{})
		err := a.OpenPullRequests(context.Background())
		assert.ErrorIs(t, err, assert.AnError)
	})
	t.Run("it should not write if there is nothing to write", func(t *testing.T) {
		// we have an error on the outputter which should not be triggered as we don't call the outputter
		a := NewApp(slogger.NewDevNullLogger(), &config.HubSphere{}, &mockManager{}, &mockOutputter{err: assert.AnError})
		err := a.OpenPullRequests(context.Background())
		assert.NoError(t, err)
	})
	t.Run("it should return an error if there are entries to write and the writer returns an error", func(t *testing.T) {
		// we have an error on the outputter which should not be triggered as we don't call the outputter
		var prs []model.RepositoryPR
		prs = append(prs, model.RepositoryPR{
			RepoName: repoName,
			Pr:       []model.PullRequest{mockPR}})

		a := NewApp(slogger.NewDevNullLogger(), &config.HubSphere{}, &mockManager{pr: prs}, &mockOutputter{err: assert.AnError})
		err := a.OpenPullRequests(context.Background())
		assert.ErrorIs(t, err, assert.AnError)
	})
	t.Run("it should send the prs to the outputter when there are no errors", func(t *testing.T) {
		var prs []model.RepositoryPR
		prs = append(prs, model.RepositoryPR{
			RepoName: repoName,
			Pr:       []model.PullRequest{mockPR}})
		prs = append(prs, model.RepositoryPR{
			RepoName: "another-example-repo",
			Pr:       []model.PullRequest{mockPR}})

		outputter := &mockOutputter{}
		a := NewApp(slogger.NewDevNullLogger(), &config.HubSphere{}, &mockManager{pr: prs}, outputter)
		assert.Len(t, outputter.entries, 0)
		err := a.OpenPullRequests(context.Background())
		assert.NoError(t, err)
		assert.Len(t, outputter.entries, 2)
	})
	t.Run("it should only repos with prs to the outputter", func(t *testing.T) {
		var prs []model.RepositoryPR
		prs = append(prs, model.RepositoryPR{
			RepoName: repoName,
			Pr:       []model.PullRequest{mockPR}})

		prs = append(prs, model.RepositoryPR{
			RepoName: "no-pull-requests-on-this-repo",
		})
		outputter := &mockOutputter{}
		a := NewApp(slogger.NewDevNullLogger(), &config.HubSphere{}, &mockManager{pr: prs}, outputter)
		assert.Len(t, outputter.entries, 0)
		err := a.OpenPullRequests(context.Background())
		assert.NoError(t, err)
		assert.Len(t, outputter.entries, 1)
	})
}

func TestAppAvailableFilters(t *testing.T) {
	t.Run("it should return an error if the outputter returns an error", func(t *testing.T) {
		cfg := &config.HubSphere{
			PullRequest: config.PullRequest{
				Filters: map[string]config.Filter{
					"test-filter": {
						Owner:        "test-owner",
						OwnerType:    config.OwnerType(config.OwnerTypeUser),
						RaisedBy:     "test-raiser",
						Labels:       []config.Label{"bug", "enhancement"},
						SummaryRegex: ".*",
					},
				},
			},
		}
		a := NewApp(slogger.NewDevNullLogger(), cfg, &mockManager{}, &mockOutputter{err: assert.AnError})
		err := a.AvailableFilters()
		assert.ErrorIs(t, err, assert.AnError)
	})
	t.Run("it should run without error", func(t *testing.T) {
		cfg := &config.HubSphere{
			PullRequest: config.PullRequest{
				Filters: map[string]config.Filter{
					"my-important-filter": {
						Owner:        "xorima",
						OwnerType:    config.OwnerType(config.OwnerTypeUser),
						RaisedBy:     "renovate",
						Labels:       []config.Label{"bug", "enhancement"},
						SummaryRegex: ".*",
					},
				},
			},
		}
		outputter := &mockOutputter{}
		a := NewApp(slogger.NewDevNullLogger(), cfg, &mockManager{}, outputter)
		err := a.AvailableFilters()
		assert.NoError(t, err)
		assert.Len(t, outputter.entries, 1)
		assert.Len(t, outputter.entries["my-important-filter"], 7)
	})
}
