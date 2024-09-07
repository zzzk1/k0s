package resource

type ApiVersion string
type Kind string
type RestartPolicy string

type Metadata struct {
	Name   string            `mapstructure:"name"`
	Labels map[string]string `mapstructure:"labels"`
}

type Spec struct {
	Containers    []Container   `mapstructure:"containers"`
	RestartPolicy RestartPolicy `mapstructure:"restartPolicy"`
}

type Container struct {
	Name            string `mapstructure:"name"`
	Image           string `mapstructure:"image"`
	ImagePullPolicy string `mapstructure:"imagePullPolicy"`
	Ports           []Port `mapstructure:"ports"`
}

type Port struct {
	Name          string `mapstructure:"name"`
	ContainerPort int32  `mapstructure:"containerPort"`
	Protocol      string `mapstructure:"protocol"`
}

type Pod struct {
	ApiVersion ApiVersion `mapstructure:"apiVersion"`
	Kind       Kind       `mapstructure:"kind"`
	Metadata   Metadata   `mapstructure:"metadata"`
	Spec       Spec       `mapstructure:"spec"`
}
