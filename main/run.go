package main

import (
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(cmdCommand)
}

var cmdCommand = &cobra.Command{
	Use: "run COMMAND",
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)

		config, err := dkr.ParseConfig(home, command)
		osexit.ExitOnError(err)
		sandboxImage, sandboxArgs := config.SandboxCommand()
		err = dkr.Sandbox(sandboxImage, sandboxArgs...)
		osexit.ExitOnError(err)
	},
}

