package main

import (
	"fmt"
	"github.com/fatih/color"
	newdkr "github.com/hekonsek/dkr/dkr"
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

		bashrc, err := newdkr.NewBashrc()
		osexit.ExitOnError(err)
		home, err := newdkr.NewDkrHome()
		osexit.ExitOnError(err)

		if !bashrc.HasPath() {
			err = bashrc.AddPath(home.Bin())
			osexit.ExitOnError(err)
			fmt.Printf("Commands directory %s was added to PATH in bashrc file. Please run the following command to reload your shell: %s\n",
				color.GreenString(home.Bin()), color.GreenString(". ~/.bashrc"))
		}

		configYml, err := newdkr.ImportConfigYml(command)
		osexit.ExitOnError(err)

		err = newdkr.SaveConfig(home, command, configYml)
		osexit.ExitOnError(err)
		err = bashrc.AddCommandProxy(home.Bin(), command)
		osexit.ExitOnError(err)
		fmt.Printf("Command %s installed.\n", color.GreenString(command))
	},
}
