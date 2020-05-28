package main

import (
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(cmdParentCommand)
}

var cmdParentCommand = &cobra.Command{
	Use:   "cmd",

	Run: func(cmd *cobra.Command, args []string) {
		osexit.ExitOnError(cmd.Help())
	},
}