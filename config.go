package dkr

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

type Config struct {
	Name       string   `yaml:"-"`
	Image      string   `yaml:"image"`
	Entrypoint []string `yaml:"entrypoint"`
}

func NewConfig(name string, image string, entrypoint []string) *Config {
	return &Config{Name: name, Image: image, Entrypoint: entrypoint}
}

func ParseConfig(home *dkrHome, command string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path.Join(home.Root, command+".yml"))
	config := Config{}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	config.Name = command
	return &config, nil
}

func (config *Config) Save(home *dkrHome) error {
	yml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return SaveConfig(home, config.Name, yml)
}

func SaveConfig(home *dkrHome, command string, configYml []byte) error {
	return ioutil.WriteFile(path.Join(home.Root, command+".yml"), configYml, 0644)
}
