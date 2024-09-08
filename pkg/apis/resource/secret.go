package resource

type SecretType string

type Secret struct {
	ApiVersion ApiVersion `mapstructure:"apiVersion"`
	Kind       Kind       `mapstructure:"kind"`
	Metadata   Metadata   `mapstructure:"metadata"`
	Type       SecretType `mapstructure:"type"`
	Data       Data       `mapstructure:"data"`
}
