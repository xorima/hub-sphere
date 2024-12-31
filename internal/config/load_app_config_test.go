package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const (
	exampleOrg      = "sous-chefs"
	exampleRaisedBy = "xorima"
)

func TestLoadAppConfig(t *testing.T) {
	t.Run("it should return an error if it cannot load the config from the user's home directory", func(t *testing.T) {
		t.Setenv("HOME", "")
		cfg, err := LoadAppConfig("")
		assert.ErrorContains(t, err, "$HOME is not defined")
		assert.Nil(t, cfg)
	})
	t.Run("it should return an error if the config file cannot be read", func(t *testing.T) {
		cfg, err := LoadAppConfig(t.TempDir())
		assert.ErrorIs(t, err, ErrReadingConfigFile)
		assert.Nil(t, cfg)
	})
	t.Run("it should return an error if the owner is empty", func(t *testing.T) {
		wantCfg := HubSphere{PullRequest: PullRequest{Filters: map[string]Filter{"foo": {
			OwnerType:    OwnerType(OwnerTypeUser),
			RaisedBy:     exampleRaisedBy,
			Labels:       nil,
			SummaryRegex: "\\.*",
			Owner:        "",
		}}}}

		_, err := LoadAppConfig(writeConfig(t, wantCfg))
		assert.ErrorIs(t, err, ErrOwnerNameMustBeDefined)

	})
	t.Run("it should return an error if the owner type is invalid", func(t *testing.T) {
		wantCfg := HubSphere{PullRequest: PullRequest{Filters: map[string]Filter{"foo": {
			OwnerType:    OwnerType(""),
			RaisedBy:     exampleRaisedBy,
			Labels:       nil,
			SummaryRegex: "\\.*",
			Owner:        exampleOrg,
		}}}}

		_, err := LoadAppConfig(writeConfig(t, wantCfg))
		assert.ErrorContains(t, err, ErrInvalidOwnerType.Error())
	})
	t.Run("it should return an error if the raised by is empty", func(t *testing.T) {
		wantCfg := HubSphere{PullRequest: PullRequest{Filters: map[string]Filter{"foo": {
			OwnerType:    OwnerType(OwnerTypeOrg),
			RaisedBy:     "",
			Labels:       nil,
			SummaryRegex: "\\.*",
			Owner:        exampleOrg,
		}}}}

		_, err := LoadAppConfig(writeConfig(t, wantCfg))
		assert.ErrorIs(t, err, ErrRaisedByMustBeDefined)
	})
	t.Run("it should return an error if the the regex is invalid", func(t *testing.T) {
		wantCfg := HubSphere{PullRequest: PullRequest{Filters: map[string]Filter{"foo": {
			OwnerType:    OwnerType(OwnerTypeOrg),
			RaisedBy:     exampleRaisedBy,
			Labels:       nil,
			SummaryRegex: "(abc",
			Owner:        exampleOrg,
		}}}}

		_, err := LoadAppConfig(writeConfig(t, wantCfg))
		assert.ErrorIs(t, err, ErrSummaryRegexInvalid)
	})
	t.Run("it should load the config file from disk without error", func(t *testing.T) {
		cfg, err := LoadAppConfig("./testdata/example.yaml")
		assert.NoError(t, err)
		assert.Len(t, cfg.PullRequest.Filters, 1)
		key := "sous-chefs-renovate"
		assert.Contains(t, cfg.PullRequest.Filters, key)
		assert.Equal(t, OwnerType(OwnerTypeOrg), cfg.PullRequest.Filters[key].OwnerType)
		assert.Equal(t, "", cfg.PullRequest.Filters[key].SummaryRegex)
		assert.Contains(t, cfg.PullRequest.Filters[key].Labels, Label("type/renovate"))
		assert.Equal(t, "renovate", cfg.PullRequest.Filters[key].RaisedBy)
		assert.Equal(t, "sous-chefs", cfg.PullRequest.Filters[key].Owner)
	})
	t.Run("it should allow passing in a custom viper", func(t *testing.T) {
		cfg, err := LoadAppConfig("./testdata/example.yaml", WithViper(viper.New()))
		assert.NoError(t, err)
		assert.Len(t, cfg.PullRequest.Filters, 1)
	})
	t.Run("it should error if no config file is found", func(t *testing.T) {
		cfg, err := LoadAppConfig("")
		assert.ErrorContains(t, err, "error reading config file")
		assert.Nil(t, cfg)
	})
}

func writeConfig(t *testing.T, cfg HubSphere) string {
	b, err := yaml.Marshal(cfg)
	assert.NoError(t, err)
	path := t.TempDir()
	p := path + "/cfg.yaml"
	assert.NoError(t, os.WriteFile(p, b, 0644))
	return p
}
