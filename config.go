package dkr

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
)

type Config struct {
	Name       string `yaml:"-"`
	Image      string `yaml:"image"`
	Entrypoint string `yaml:"entrypoint"`
}

func (config *Config) Save(home *dkrHome) error {
	yml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(home.Root, config.Name + ".yml"), yml, 0644)
	return err
}