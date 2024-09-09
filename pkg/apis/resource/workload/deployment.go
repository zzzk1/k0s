package workload

import "k0s/pkg/apis/resource/common"

type Deployment struct {
	ApiVersion common.ApiVersion `mapstructure:"apiVersion"`
	Kind       common.Kind       `mapstructure:"kind"`
	Metadata   common.ObjectMeta `mapstructure:"metadata"`
	Spec       DeploymentSpec    `mapstructure:"spec"`
}

type DeploymentSpec struct {
	Replicas int32                `mapstructure:"replicas"`
	Selector common.LabelSelector `mapstructure:"selector"`
	Template PodTemplateSpec      `mapstructure:"template"`
}
