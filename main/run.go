package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	DkrCommand.AddCommand(cmdCommand)
}

var cmdCommand = &cobra.Command{
	Use:                "run COMMAND",
	Short:              "Executes installed command",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)

		config, err := dkr.ParseConfig(home, command)
		osexit.ExitOnError(err)
		if config == nil {
			osexit.ExitBecauseError(fmt.Sprintf(
				"Command %s not installed. Have you tried installing that command by executing %s ?",
				color.GreenString(command), color.GreenString("dkr cmd install "+command)))
		}
		err = dkr.Sandbox(config.Image, config.Entrypoint, args[1:], nil)
		osexit.ExitOnError(err)
	},
}
