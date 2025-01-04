package data

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/google/go-github/v68/github"
	"golang.org/x/oauth2"
)

const tokenEnvVar = "GITHUB_TOKEN"

var ErrNoGithubToken = errors.New("the env var 'GITHUB_TOKEN' does not exist")

type Repository = github.Repository

type GithubClient struct {
	client *github.Client
	log    *slog.Logger
}
type GithubClientOpts struct {
	log       *slog.Logger
	envKey    string
	transport http.RoundTripper
}
type OptsFunc func(o *GithubClientOpts)

func newGithubClientOpts(log *slog.Logger) *GithubClientOpts {
	return &GithubClientOpts{
		log:    log,
		envKey: tokenEnvVar,
	}
}

func WithTransport(transport http.RoundTripper) OptsFunc {
	return func(o *GithubClientOpts) {
		o.transport = transport
	}
}

func (o *GithubClientOpts) getHttpClient() (*http.Client, error) {
	client := http.DefaultClient
	if o.transport != nil {
		client.Transport = o.transport
		return client, nil
	}
	token := os.Getenv(o.envKey)
	if token == "" {
		o.log.Error("the github token is not set in the environment variables", slog.String("key", o.envKey))
		return nil, ErrNoGithubToken
	}
	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}))
	client.Transport = tc.Transport
	return client, nil
}

func NewGithubClient(ctx context.Context, log *slog.Logger, clientOpts ...OptsFunc) (*GithubClient, error) {
	opts := newGithubClientOpts(log)
	for _, o := range clientOpts {
		o(opts)
	}
	c, err := opts.getHttpClient()
	if err != nil {
		return nil, err
	}
	return &GithubClient{client: github.NewClient(c), log: log}, nil
}
