package main

import (
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	DkrCommand.AddCommand(sandboxCommand)
}

var sandboxCommand = &cobra.Command{
	Use:   "sandbox IMAGE [args...]",
	Short: "Executes given command in a dockerized sandbox",
	Run: func(cmd *cobra.Command, args []string) {
		err := dkr.Sandbox(args[0], nil, args[1:], nil)
		osexit.ExitOnError(err)
	},
}
