package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/dkr"
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

		configYml, err := dkr.ImportConfigYml(command)
		osexit.ExitOnError(err)

		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)
		err = dkr.SaveConfig(home, command, configYml)
		osexit.ExitOnError(err)
		fmt.Printf("Command %s installed.\n", color.GreenString(command))

		bashrc, err := newdkr.NewBashrc()
		osexit.ExitOnError(err)
		err = bashrc.AddAlias(command)
		osexit.ExitOnError(err)
		fmt.Printf("Bash alias for command %s was added to %s file. Please run the following command to reload your shell: %s\n",
			color.GreenString(command), color.GreenString("~/.bashrc"), color.GreenString(". ~/.bashrc"))
	},
}
