package common

type ApiVersion string
type Kind string

type ObjectMeta struct {
	Name   string            `mapstructure:"name"`
	Labels map[string]string `mapstructure:"labels"`
}
