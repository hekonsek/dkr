package dkr

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const bashrcTagDkr = "#dkr"

type Bashrc struct {
	filepath string
	lines    []string
}

func NewBashrc() (*Bashrc, error) {
	bashrc := &Bashrc{}
	err := bashrc.load()
	if err != nil {
		return nil, err
	}
	return bashrc, nil
}

func (bashrc *Bashrc) load() error {
	var home, err = os.UserHomeDir()
	if err != nil {
		return err
	}
	bashrc.filepath = path.Join(home, ".bashrc")
	bashrcBytes, err := ioutil.ReadFile(bashrc.filepath)
	if err != nil {
		return err
	}
	bashrc.lines = strings.Split(string(bashrcBytes), "\n")
	return nil
}

func (bashrc *Bashrc) Lines() []string {
	return bashrc.lines
}

func (bashrc *Bashrc) DkrLines() []string {
	var dkrLines []string
	for _, line := range bashrc.lines {
		if strings.HasSuffix(line, bashrcTagDkr) {
			dkrLines = append(dkrLines, line)
		}
	}
	return dkrLines
}

func (bashrc *Bashrc) HasAlias(command string) (bool, error) {
	var aliasMatcher, err = regexp.Compile(fmt.Sprintf("alias %s='.*", command))
	if err != nil {
		return false, err
	}
	for _, line := range bashrc.DkrLines() {
		if aliasMatcher.MatchString(line) {
			return true, nil
		}
	}
	return false, nil
}

func (bashrc *Bashrc) AddAlias(command string) error {
	hasAlias, err := bashrc.HasAlias(command)
	if err != nil {
		return err
	}
	if !hasAlias {
		var alias = fmt.Sprintf("alias %s='dkr run %s' %s", command, command, bashrcTagDkr)
		bashrc.lines = append(bashrc.lines, alias)
		return ioutil.WriteFile(bashrc.filepath, []byte(strings.Join(bashrc.lines, "\n")), 0644)
	}
	return nil
}
