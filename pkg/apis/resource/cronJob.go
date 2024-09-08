package resource

type Schedule string

type CronJobSpec struct {
	Schedule    Schedule    `mapstructure:"schedule"`
	JobTemplate JobTemplate `mapstructure:"jobTemplate"`
}

type JobTemplate struct {
	Spec JobTemplateSpec `mapstructure:"spec"`
}

type JobTemplateSpec struct {
	Template Template `mapstructure:"template"`
}

type CronJob struct {
	ApiVersion ApiVersion  `mapstructure:"apiVersion"`
	Kind       Kind        `mapstructure:"kind"`
	Metadata   Metadata    `mapstructure:"metadata"`
	Spec       CronJobSpec `mapstructure:"spec"`
}
