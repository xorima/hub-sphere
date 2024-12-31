package config

import "errors"

type PullRequest struct {
	Filters map[string]Filter `yaml:"filters" mapstructure:"filters"`
}

func (p *PullRequest) validate() error {
	var errs []error
	for _, v := range p.Filters {
		err := v.validate()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
