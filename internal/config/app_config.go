package config

type HubSphere struct {
	PullRequest PullRequest `yaml:"pullRequest" mapstructure:"pullRequest"`
}

func (h *HubSphere) validate() error {
	return h.PullRequest.validate()
}
