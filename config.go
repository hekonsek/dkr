package dkr

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"strings"
)

type Config struct {
	Name       string `yaml:"-"`
	Image      string `yaml:"image"`
	Entrypoint string `yaml:"entrypoint"`
}

func ParseConfig(home *dkrHome, command string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path.Join(home.Root, command + ".yml"))
	config := Config{}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (config *Config) Save(home *dkrHome) error {
	yml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(home.Root, config.Name + ".yml"), yml, 0644)
	return err
}

func (config *Config) SandboxCommand() (string, []string) {
	return config.Image, strings.Split(config.Entrypoint, " ")
}