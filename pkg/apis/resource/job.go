package resource

type JobSpec struct {
	Template     Template     `mapstructure:"template"`
	BackoffLimit BackoffLimit `mapstructure:"backoffLimit"`
}

type Job struct {
	ApiVersion ApiVersion `mapstructure:"apiVersion"`
	Kind       Kind       `mapstructure:"kind"`
	Metadata   Metadata   `mapstructure:"metadata"`
	Spec       JobSpec    `mapstructure:"spec"`
}
