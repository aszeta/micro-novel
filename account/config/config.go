package config

import (
	"io/ioutil"

	"sigs.k8s.io/yaml"
)

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}

func ProvideConfig(configPath string) *Config {
	conf := Config{}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
