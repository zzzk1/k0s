package runtime

type ContainerPort struct {
	Name          string `mapstructure:"name"`
	ContainerPort int32  `mapstructure:"containerPort"`
	Protocol      string `mapstructure:"protocol"`
}

type Container struct {
	Name            string          `mapstructure:"name"`
	Image           string          `mapstructure:"image"`
	ImagePullPolicy string          `mapstructure:"imagePullPolicy"`
	Args            []string        `mapstructure:"args"`
	Command         []string        `mapstructure:"command"`
	Ports           []ContainerPort `mapstructure:"ports"`
}
