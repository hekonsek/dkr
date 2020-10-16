package dkr

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const bashrcTagDkr = "#dkr"

type Bashrc struct {
	filepath string
	lines    []string
}

func NewBashrcFromFile(path string) (*Bashrc, error) {
	bashrc := &Bashrc{filepath: path}
	err := bashrc.load()
	if err != nil {
		return nil, err
	}
	return bashrc, nil
}

func NewBashrc() (*Bashrc, error) {
	var home, err = os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return NewBashrcFromFile(path.Join(home, ".bashrc"))
}

func (bashrc *Bashrc) load() error {
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

func (bashrc *Bashrc) HasPath() bool {
	for _, line := range bashrc.DkrLines() {
		if strings.HasPrefix(line, "export PATH=") {
			return true
		}
	}
	return false
}

func (bashrc *Bashrc) AddPath(binPath string) error {
	if !bashrc.HasPath() {
		var dkrPathLine = fmt.Sprintf("export PATH=${PATH}:%s %s", binPath, bashrcTagDkr)
		bashrc.lines = append(bashrc.lines, dkrPathLine)
		return ioutil.WriteFile(bashrc.filepath, []byte(strings.Join(bashrc.lines, "\n")), 0644)
	}
	return bashrc.load()
}

func (bashrc *Bashrc) AddCommandProxy(binPath string, command string) error {
	if !bashrc.HasPath() {
		err := bashrc.AddPath(binPath)
		if err != nil {
			return err
		}
	}

	proxyPath := path.Join(binPath, command)
	script := fmt.Sprintf("#!/bin/bash\ndkr run %s \"$@\"\n", command)
	return ioutil.WriteFile(proxyPath, []byte(script), 0755)
}
