package config

import (
	"k0s/pkg/apis/resource"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	t.Run("pod", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/test-pod.yaml")
		actual := (*apiResource).(*resource.Pod)
		expected := resource.Pod{
			ApiVersion: resource.ApiVersion("v1"),
			Kind:       resource.Kind("Pod"),
			Metadata: resource.Metadata{
				Name: "my-pod",
				Labels: map[string]string{
					"app": "my-application",
				},
			},
			Spec: resource.PodSpec{
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

	t.Run("Job", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/test-job.yaml")
		actual := *(*apiResource).(*resource.Job)

		expected := resource.Job{
			ApiVersion: resource.ApiVersion("batch/v1"),
			Kind:       resource.Kind("Job"),
			Metadata: resource.Metadata{
				Name: "echo-hello",
				Labels: map[string]string{
					"task": "once",
				},
			},
			Spec: resource.JobSpec{
				Template: resource.Template{
					Spec: resource.Spec{
						Containers: []resource.Container{
							{
								Name:            "echo-hello",
								Image:           "python:latest",
								ImagePullPolicy: "IfNotPresent",
								Command:         []string{"python", "-c"},
								Args:            []string{"print('Hello from the Kubernetes job')"},
							},
						},
						RestartPolicy: "Never",
					},
				},
				BackoffLimit: 4,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("CronJob", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/test-cronjob.yaml")
		actual := *(*apiResource).(*resource.CronJob)
		expected := resource.CronJob{
			ApiVersion: resource.ApiVersion("batch/v1"),
			Kind:       resource.Kind("CronJob"),
			Metadata:   resource.Metadata{Name: "echo-hello-10s"},
			Spec: resource.CronJobSpec{
				Schedule: "* * * * 10",
				JobTemplate: resource.JobTemplate{
					Spec: resource.JobTemplateSpec{
						resource.Template{
							Spec: resource.Spec{
								Containers: []resource.Container{
									{Name: "echo-hello-10s", Image: "busybox:latest", ImagePullPolicy: "IfNotPresent", Command: []string{"/bin/sh", "-c", "date; echo Hello!"}},
								},
								RestartPolicy: "OnFailure",
							},
						},
					},
				},
			},
		}
		assert.Equal(t, expected, actual)
	})
}
