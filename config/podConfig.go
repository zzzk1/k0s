package config

import (
	"k0s/pkg/apis/resource"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string) *resource.Pod {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading pod config file: %v", err)
		return nil
	}

	var pod resource.Pod
	if err := viper.Unmarshal(&pod); err != nil {
		log.Printf("Error unmarshaling pod config file: %v", err)
		return nil
	}

	return &pod
}
