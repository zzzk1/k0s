package resource

type ConfigMap struct {
	ApiVersion ApiVersion `mapstructure:"apiVersion"`
	Kind       Kind       `mapstructure:"kind"`
	Metadata   Metadata   `mapstructure:"metadata"`
	Data       Data       `mapstructure:"data"`
}
