package resource

type Replicas int

type Selector struct {
	MatchLabels map[string]string
}

type DeploymentTemplate struct {
	Metadata Metadata `mapstructure:"metadata"`
	Spec     Spec     `mapstructure:"spec"`
}

type DeploymentSpec struct {
	Replicas Replicas           `mapstructure:"replicas"`
	Selector Selector           `mapstructure:"selector"`
	Template DeploymentTemplate `mapstructure:"template"`
}

type Deployment struct {
	ApiVersion ApiVersion     `mapstructure:"apiVersion"`
	Kind       Kind           `mapstructure:"kind"`
	Metadata   Metadata       `mapstructure:"metadata"`
	Spec       DeploymentSpec `mapstructure:"spec"`
}
