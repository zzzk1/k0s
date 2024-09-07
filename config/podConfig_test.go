package config

import (
	"k0s/pkg/apis/resource"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("pod", func(t *testing.T) {
		actual := LoadConfig("../config/template/test-pod.yaml")

		expected := resource.Pod{
			ApiVersion: resource.ApiVersion("v1"),
			Kind:       resource.Kind("Pod"),
			Metadata: resource.Metadata{
				Name: "my-pod",
				Labels: map[string]string{
					"app": "my-application",
				},
			},
			Spec: resource.Spec{
				Containers: []resource.Container{
					{
						Name:            "nginx-container",
						Image:           "nginx:latest",
						ImagePullPolicy: "IfNotPresent",
						Ports: []resource.Port{
							{
								Name:          "http",
								ContainerPort: 80,
								Protocol:      "TCP",
							},
						},
					},
				},
				RestartPolicy: "Always",
			},
		}
		assert.Equal(t, expected, *actual)
	})
}
