package storage

import "k0s/pkg/apis/resource/common"

type Secret struct {
	ApiVersion common.ApiVersion `mapstructure:"apiVersion"`
	Kind       common.Kind       `mapstructure:"kind"`
	Metadata   common.ObjectMeta `mapstructure:"metadata"`
	Type       string            `mapstructure:"type"`
	Data       map[string]string `mapstructure:"data"`
}
