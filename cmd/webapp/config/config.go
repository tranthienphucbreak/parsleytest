package config

import (
	"gopkg.in/yaml.v3"
	//"fmt"
	"os"
)

type ServiceConfig struct {
	Port int64 `yaml:"port"`
}

func ReadConfig() *ServiceConfig {
	yamlFile, err := os.Open("./settings.yaml")
	if err != nil {
		yamlFile, err = os.Open("./../../settings.yaml")
		if err != nil {
			panic(err)
		}
	}
	config := ServiceConfig{}
	if yamlFile != nil {
		decoder := yaml.NewDecoder(yamlFile)
		if err := decoder.Decode(&config); err != nil {
			panic(err)
		}
	}
	return &config
}
