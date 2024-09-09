package resource

type RestartPolicy string
type BackoffLimit int64

type Data map[string]interface{}

type Template struct {
	Spec Spec `mapstructure:",spec"`
}
type Spec struct {
	Containers    []Container   `mapstructure:"containers"`
	RestartPolicy RestartPolicy `mapstructure:"restartPolicy"`
}

type Container struct {
	Name            string   `mapstructure:"name"`
	Image           string   `mapstructure:"image"`
	ImagePullPolicy string   `mapstructure:"imagePullPolicy"`
	Command         []string `mapstructure:"command"`
	Args            []string `mapstructure:"args"`
	Ports           []Port   `mapstructure:"ports"`
}

type Port struct {
	Name          string `mapstructure:"name"`
	ContainerPort int32  `mapstructure:"containerPort"`
	Protocol      string `mapstructure:"protocol"`
}
