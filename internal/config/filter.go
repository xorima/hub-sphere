package config

import (
	"errors"
	"regexp"
)

var (
	ErrOwnerNameMustBeDefined = errors.New("owner name is empty, it must be defined")
	ErrRaisedByMustBeDefined  = errors.New("raised by is empty, it must be defined")
	ErrSummaryRegexInvalid    = errors.New("the given summary regex is invalid")
)

type Filter struct {
	Owner             string       `yaml:"owner" mapstructure:"owner"`
	OwnerType         OwnerType    `yaml:"type" mapstructure:"type"`
	RaisedBy          string       `yaml:"raisedBy" mapstructure:"raisedBy"`
	Labels            []Label      `yaml:"labels" mapstructure:"labels"`
	SummaryRegex      string       `yaml:"summaryRegex" mapstructure:"summaryRegex"`
	Repositories      []Repository `yaml:"repositories" mapstructure:"repositories"`
	summaryRegexCache *regexp.Regexp
}

func (f *Filter) validate() error {
	var errs []error
	if f.Owner == "" {
		errs = append(errs, ErrOwnerNameMustBeDefined)
	}
	errs = append(errs, f.OwnerType.validate())
	if f.RaisedBy == "" {
		errs = append(errs, ErrRaisedByMustBeDefined)
	}
	r, err := regexp.Compile(f.SummaryRegex)
	if err != nil {
		errs = append(errs, errors.Join(ErrSummaryRegexInvalid, err))
	}
	f.summaryRegexCache = r

	return errors.Join(errs...)
}
