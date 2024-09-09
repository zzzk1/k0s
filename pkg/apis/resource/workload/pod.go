package workload

import (
	"k0s/pkg/apis/resource/common"
	"k0s/pkg/apis/resource/runtime"
)

type Pod struct {
	ApiVersion common.ApiVersion `mapstructure:"apiVersion"`
	Kind       common.Kind       `mapstructure:"kind"`
	Metadata   common.ObjectMeta `mapstructure:"metadata"`
	Spec       PodSpec           `mapstructure:"spec"`
	Container  runtime.Container `mapstructure:"container"`
}

type PodTemplateSpec struct {
	Metadata common.ObjectMeta `mapstructure:"metadata"`
	Spec     PodSpec           `mapstructure:"spec"`
}

type RestartPolicy string

type PodSpec struct {
	Containers    []runtime.Container `mapstructure:"containers"`
	RestartPolicy RestartPolicy       `mapstructure:"restartPolicy"`
}
