package dkr

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
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

func ImportConfigYml(command string) ([]byte, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/hekonsek/dkr-%s/master/config.yml",
		command)
	resp, err := http.Get(url)
	osexit.ExitOnError(err)
	if resp.StatusCode == 404 { // Deprecated: backward compatibility for commands in main dkr project repository
		url = fmt.Sprintf("https://raw.githubusercontent.com/hekonsek/dkr/master/commands/%s/config.yml",
			command)
		resp, err = http.Get(url)
		if err != nil {
			return nil, err
		}
	}
	if resp.StatusCode != 200 {
		osexit.ExitBecauseError("No such command.")
	}
	defer resp.Body.Close()
	config, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return config, nil
}
