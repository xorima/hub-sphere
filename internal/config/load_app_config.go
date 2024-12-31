package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

var (
	ErrReadingConfigFile    = errors.New("error reading config file")
	ErrUnableToDecodeConfig = errors.New("error decoding the configuration file")
)

type AppConfigOpts struct {
	viper *viper.Viper
}

func newAppConfigOpts() *AppConfigOpts {
	return &AppConfigOpts{
		viper: viper.New(),
	}
}

type AppConfigOptsFunc func(opts *AppConfigOpts)

func WithViper(viper *viper.Viper) func(opts *AppConfigOpts) {
	return func(opts *AppConfigOpts) {
		opts.viper = viper
	}
}

func LoadAppConfig(cfgFile string, opts ...AppConfigOptsFunc) (*HubSphere, error) {

	cfgOpts := newAppConfigOpts()
	for _, o := range opts {
		o(cfgOpts)
	}

	cfgOpts.viper.SetConfigType("yaml")

	if cfgFile != "" {
		cfgOpts.viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		cfgOpts.viper.AddConfigPath(home)
		cfgOpts.viper.SetConfigName(".hub-sphere")
	}

	if err := cfgOpts.viper.ReadInConfig(); err != nil {
		return nil, errors.Join(ErrReadingConfigFile, err)
	}
	var cfg HubSphere
	if err := cfgOpts.viper.Unmarshal(&cfg, ownerDecodeHookFunc()); err != nil {
		return nil, errors.Join(ErrUnableToDecodeConfig, err)
	}
	return &cfg, cfg.validate()
}
