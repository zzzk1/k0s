package workload

import "k0s/pkg/apis/resource/common"

type CronJob struct {
	ApiVersion common.ApiVersion `mapstructure:"apiVersion"`
	Kind       common.Kind       `mapstructure:"kind"`
	Metadata   common.ObjectMeta `mapstructure:"metadata"`
	Spec       CronJobSpec       `mapstructure:"spec"`
}

type CronJobSpec struct {
	JobTemplate JobTemplate `mapstructure:"jobTemplate"`
	Schedule    string      `mapstructure:"schedule"`
}

type JobTemplate struct {
	Metadata common.ObjectMeta `mapstructure:"metadata"`
	Spec     JobSpec           `mapstructure:"spec"`
}
