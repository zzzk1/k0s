package resource

type ApiVersion string
type Kind string

type Metadata struct {
	Name   string            `mapstructure:"name"`
	Labels map[string]string `mapstructure:"labels"`
}
