package config

import (
	"k0s/pkg/apis/resource/common"
	"k0s/pkg/apis/resource/runtime"
	"k0s/pkg/apis/resource/storage"
	"k0s/pkg/apis/resource/workload"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	t.Run("pod", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/simple-pod.yaml")
		actual := (*apiResource).(workload.Pod)
		expected := workload.Pod{
			ApiVersion: common.ApiVersion("v1"),
			Kind:       common.Kind("Pod"),
			Metadata: common.ObjectMeta{
				Name: "my-pod",
				Labels: map[string]string{
					"app": "my-application",
				},
			},
			Spec: workload.PodSpec{
				Containers: []runtime.Container{
					{
						Name:            "nginx-container",
						Image:           "nginx:latest",
						ImagePullPolicy: "IfNotPresent",
						Ports: []runtime.ContainerPort{
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
		assert.Equal(t, expected, actual)
	})

	t.Run("Job", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/simple-job.yaml")
		actual := *(*apiResource).(*workload.Job)

		expected := workload.Job{
			ApiVersion: common.ApiVersion("batch/v1"),
			Kind:       common.Kind("Job"),
			Metadata: common.ObjectMeta{
				Name: "echo-hello",
				Labels: map[string]string{
					"task": "once",
				},
			},
			Spec: workload.JobSpec{
				Template: workload.PodTemplateSpec{
					Spec: workload.PodSpec{
						Containers: []runtime.Container{
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
		apiResource := LoadConfig("../config/template/simple-cronjob.yaml")
		actual := *(*apiResource).(*workload.CronJob)
		expected := workload.CronJob{
			ApiVersion: common.ApiVersion("batch/v1"),
			Kind:       common.Kind("CronJob"),
			Metadata:   common.ObjectMeta{Name: "echo-hello-10s"},
			Spec: workload.CronJobSpec{
				Schedule: "* * * * 10",
				JobTemplate: workload.JobTemplate{
					Spec: workload.JobSpec{
						Template: workload.PodTemplateSpec{
							Spec: workload.PodSpec{
								Containers: []runtime.Container{
									{
										Name:            "echo-hello-10s",
										Image:           "busybox:latest",
										ImagePullPolicy: "IfNotPresent",
										Command:         []string{"/bin/sh", "-c", "date; echo Hello!"},
									},
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

	t.Run("ConfigMap", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/simple-configmap.yaml")
		actual := *(*apiResource).(*storage.ConfigMap)
		expected := storage.ConfigMap{
			ApiVersion: common.ApiVersion("v1"),
			Kind:       common.Kind("ConfigMap"),
			Metadata: common.ObjectMeta{
				Name: "mysql-config",
			},
			Data: map[string]string{
				"host": "db.dev.local",
				"port": "3306"},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("Secret", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/simple-secret.yaml")
		actual := *(*apiResource).(*storage.Secret)
		expected := storage.Secret{
			ApiVersion: common.ApiVersion("v1"),
			Kind:       common.Kind("Secret"),
			Metadata: common.ObjectMeta{
				Name: "user-basic-auth",
			},
			Type: "kubernetes.io/basic-auth",
			Data: map[string]string{
				"username": "cm9vdA==",
				"password": "MTIzNDU2",
			},
		}

		assert.Equal(t, expected, actual)
	})

	t.Run("Deployment", func(t *testing.T) {
		apiResource := LoadConfig("../config/template/simple-deploy.yaml")
		actual := *(*apiResource).(*workload.Deployment)
		expected := workload.Deployment{
			ApiVersion: "apps/v1",
			Kind:       "Deployment",
			Metadata: common.ObjectMeta{
				Name: "nginx-deploy",
				Labels: map[string]string{
					"app": "nginx-deploy",
				},
			},
			Spec: workload.DeploymentSpec{
				Replicas: 1,
				Selector: common.LabelSelector{
					MatchLabels: map[string]string{
						"app": "my-nginx",
					},
				},
				Template: workload.PodTemplateSpec{
					Metadata: common.ObjectMeta{
						Name: "my-nginx",
						Labels: map[string]string{
							"app": "my-nginx",
						},
					},
					Spec: workload.PodSpec{
						Containers: []runtime.Container{
							{
								Name:            "my-nginx",
								Image:           "nginx:latest",
								ImagePullPolicy: "IfNotPresent",
								Ports: []runtime.ContainerPort{
									{
										ContainerPort: 80,
										Protocol:      "TCP",
									},
								},
							},
						},
						RestartPolicy: "Always",
					},
				},
			},
		}

		assert.Equal(t, expected, actual)
	})
}
