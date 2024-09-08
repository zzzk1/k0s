package resource

type PodSpec struct {
	Containers    []Container   `mapstructure:"containers"`
	RestartPolicy RestartPolicy `mapstructure:"restartPolicy"`
}

type Pod struct {
	ApiVersion ApiVersion `mapstructure:"apiVersion"`
	Kind       Kind       `mapstructure:"kind"`
	Metadata   Metadata   `mapstructure:"metadata"`
	Spec       PodSpec    `mapstructure:"spec"`
}
