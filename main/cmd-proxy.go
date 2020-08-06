package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	cmdParentCommand.AddCommand(cmdProxyCommand)
}

var cmdProxyCommand = &cobra.Command{
	Use: "proxy COMMAND",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			osexit.ExitOnError(cmd.Help())
			return
		}
		command := args[0]

		target, err := dkr.CopyProxy(command)
		osexit.ExitOnError(err)
		fmt.Printf("Proxy file for command %s created: %s\n",
			color.GreenString(command), color.GreenString(target))
	},
}
