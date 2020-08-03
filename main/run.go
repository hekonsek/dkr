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
	Use:                "run COMMAND",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)

		config, err := dkr.ParseConfig(home, command)
		osexit.ExitOnError(err)
		err = dkr.Sandbox(config.Image, config.Entrypoint, args[1:]...)
		osexit.ExitOnError(err)
	},
}
