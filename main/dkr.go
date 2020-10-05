package main

import (
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

var DkrCommand = &cobra.Command{
	Use:   "dkr",
	Short: `DKR (pronounced "dockerizer") is a simple toolkit to help you dockerize your shell commands.`,

	Run: func(cmd *cobra.Command, args []string) {
		osexit.ExitOnError(cmd.Help())
	},
}

func main() {
	osexit.ExitOnError(DkrCommand.Execute())
}
