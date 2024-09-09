package common

type LabelSelector struct {
	MatchLabels map[string]string `mapstructure:"matchLabels"`
}
