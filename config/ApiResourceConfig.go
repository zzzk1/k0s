package config

import (
	"github.com/spf13/viper"
	"k0s/pkg/apis/resource/storage"
	"k0s/pkg/apis/resource/workload"
	"log"
)

type ApiResourceFactory interface {
}

func getApiResource(kind string) ApiResourceFactory {
	switch kind {
	case "Pod":
		return workload.Pod{}
	case "Job":
		return &workload.Job{}
	case "CronJob":
		return &workload.CronJob{}
	case "ConfigMap":
		return &storage.ConfigMap{}
	case "Secret":
		return &storage.Secret{}
	case "Deployment":
		return &workload.Deployment{}
	default:
		return nil
	}
}

func LoadConfig(path string) *ApiResourceFactory {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading pod config file: %v", err)
		return nil
	}

	apiResource := getApiResource(viper.GetString("kind"))
	if apiResource == nil {
		log.Printf("Unsupported kind: %v", viper.GetString("kind"))
		return nil
	}

	if err := viper.Unmarshal(&apiResource); err != nil {
		log.Printf("Error unmarshaling pod config file: %v", err)
		return nil
	}

	return &apiResource
}
