package workload

import "k0s/pkg/apis/resource/common"

type Job struct {
	ApiVersion common.ApiVersion `mapstructure:"apiVersion"`
	Kind       common.Kind       `mapstructure:"kind"`
	Metadata   common.ObjectMeta `mapstructure:"metadata"`
	Spec       JobSpec           `mapstructure:"spec"`
}

type JobSpec struct {
	Template     PodTemplateSpec `mapstructure:"template"`
	BackoffLimit int32           `mapstructure:"backoffLimit"`
}
