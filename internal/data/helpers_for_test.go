package data

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xorima/slogger"
	"golang.org/x/oauth2"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

var (
	AccessToken = "TEST_TOKEN_HERE" // Update this when creating new tests
)

type mockTransport struct {
}

func (m *mockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, assert.AnError
}

func newVcrGithubClient(t *testing.T, vcrPath, token string) (*GithubClient, *recorder.Recorder, error) {
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	r, err := recorder.New(vcrPath, recorder.WithRealTransport(tc.Transport), recorder.WithMode(recorder.ModeReplayWithNewEpisodes), recorder.WithSkipRequestLatency(true))
	if err != nil {
		return nil, nil, err
	}
	httpClient := http.DefaultClient
	httpClient.Transport = r
	c, err := NewGithubClient(ctx, slogger.NewDevNullLogger(), WithTransport(r))
	assert.NoError(t, err)
	return c, r, nil
}
