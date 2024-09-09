package storage

import "k0s/pkg/apis/resource/common"

type ConfigMap struct {
	ApiVersion common.ApiVersion `mapstructure:"apiVersion"`
	Kind       common.Kind       `mapstructure:"kind"`
	Metadata   common.ObjectMeta `mapstructure:"metadata"`
	Data       map[string]string `mapstructure:"data"`
}
