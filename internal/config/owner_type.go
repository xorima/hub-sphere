package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type OwnerType string

var (
	OwnerTypeOrg        = "org"
	OwnerTypeUser       = "user"
	ErrInvalidOwnerType = fmt.Errorf("invalid owner type, only '%s' and '%s' are allowed", OwnerTypeOrg, OwnerTypeUser)
)

func (o *OwnerType) String() string {
	return string(*o)
}

func (o *OwnerType) validate() error {
	ot := strings.ToLower(o.String())
	if ot != OwnerTypeOrg && ot != OwnerTypeUser {
		return ErrInvalidOwnerType
	}
	return nil
}

func ownerDecodeHookFunc() viper.DecoderConfigOption {
	return func(config *mapstructure.DecoderConfig) {
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			func(
				f reflect.Type,
				t reflect.Type,
				data interface{},
			) (interface{}, error) {
				if f.Kind() != reflect.String {
					return data, nil
				}
				if t != reflect.TypeOf(OwnerType("")) {
					return data, nil
				}
				ownerType := OwnerType(strings.ToLower(data.(string)))
				if err := ownerType.validate(); err != nil {
					return nil, err
				}
				return ownerType, nil
			},
		)
	}
}
