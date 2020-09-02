package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

func init() {
	cmdParentCommand.AddCommand(cmdInstallCommand)
}

var cmdInstallCommand = &cobra.Command{
	Use: "install COMMAND",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			osexit.ExitOnError(cmd.Help())
			return
		}
		command := args[0]

		configYml, err := dkr.ImportConfigYml(command)
		osexit.ExitOnError(err)

		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)
		err = dkr.SaveConfig(home, command, configYml)
		osexit.ExitOnError(err)
		fmt.Printf("Command %s added.\n", color.GreenString(command))

		target, err := dkr.CopyProxy(command)
		osexit.ExitOnError(err)
		fmt.Printf("Proxy file for command %s created: %s\n",
			color.GreenString(command), color.GreenString(target))
	},
}
