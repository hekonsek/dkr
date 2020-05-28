package main

import (
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use: "sandbox IMAGE [args...]",
	Run: func(cmd *cobra.Command, args []string) {
		err := dkr.Sandbox(args[0], args[1:]...)
		osexit.ExitOnError(err)
	},
}