package main

import (
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func init() {
	RootCommand.AddCommand(cmdCommand)
}

var cmdCommand = &cobra.Command{
	Use: "run COMMAND",
	Run: func(cmd *cobra.Command, args []string) {
		execName := args[0]

		home, err := dkr.NewDcmHome()
		osexit.ExitOnError(err)

		configJson, err := ioutil.ReadFile(path.Join(home.Root, execName + ".yml"))
		osexit.ExitOnError(err)
		var config map[string]string
		err = yaml.Unmarshal(configJson, &config)
		osexit.ExitOnError(err)

		command := []string{"run"}
		command = append(command, config["image"])
		if config["entrypoint"] != "" {
			entryPoint := strings.Split(config["entrypoint"], " ")
			command = append(command, entryPoint...)
		}
		command = append(command, os.Args[3:]...)
		cmdx := exec.Command("dcm", command...)
		cmdx.Stdin = os.Stdin
		cmdx.Stdout = os.Stdout
		cmdx.Stderr = os.Stderr
		err = cmdx.Run()
		osexit.ExitOnError(err)
	},
}

