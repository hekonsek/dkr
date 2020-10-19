package dkr

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
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
	configBytes, err := ioutil.ReadFile(path.Join(home.Root(), command+".yml"))
	if err != nil {
		return nil, nil
	}
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
	return ioutil.WriteFile(path.Join(home.Root(), command+".yml"), configYml, 0644)
}

var NoSuchCommandError = errors.New("no such command")

func ImportConfigYml(command string) ([]byte, error) {
	for _, configUrl := range generatePossibleConfigUrls(command) {
		resp, err := http.Get(configUrl)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			continue
		}
		defer resp.Body.Close()
		config, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return config, nil
	}
	return nil, NoSuchCommandError
}

func generatePossibleConfigUrls(command string) []string {
	user := ""
	repo := ""
	canonical := ParseCanonicalCommand(command)
	if canonical != nil {
		user = canonical.User
		repo = canonical.Repo
	} else {
		user = "hekonsek"
		repo = "dkr-" + command
	}

	urlTemplate := "https://raw.githubusercontent.com/%s/%s/%s/%s.yml"
	urls := []string{}
	for _, branch := range []string{"main", "master"} {
		for _, config := range []string{"config", "dkr"} {
			urls = append(urls, fmt.Sprintf(urlTemplate, user, repo, branch, config))
		}
	}
	return urls
}

type CanonicalCommand struct {
	Repo string
	User string
}

func ParseCanonicalCommand(command string) *CanonicalCommand {
	if !IsCanonical(command) {
		return nil
	}
	commandWithoutGithubPrefix := strings.Replace(command, "github.com/", "", 1)
	commandParts := strings.Split(commandWithoutGithubPrefix, "/")
	return &CanonicalCommand{User: commandParts[0], Repo: commandParts[1]}
}

func IsCanonical(command string) bool {
	return strings.HasPrefix(command, "github.com/")
}
