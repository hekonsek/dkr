package main

import (
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	DkrCommand.AddCommand(cmdParentCommand)
}

var cmdParentCommand = &cobra.Command{
	Use:   "cmd",
	Short: "Manages commands installed on this machine",

	Run: func(cmd *cobra.Command, args []string) {
		osexit.ExitOnError(cmd.Help())
	},
}
