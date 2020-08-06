package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

func init() {
	cmdParentCommand.AddCommand(cmdAddCommand)
}

var cmdAddCommand = &cobra.Command{
	Use: "add COMMAND",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			osexit.ExitOnError(cmd.Help())
			return
		}
		command := args[0]

		url := fmt.Sprintf("https://raw.githubusercontent.com/hekonsek/dkr/master/commands/%s/config.yml",
			command)
		resp, err := http.Get(url)
		osexit.ExitOnError(err)
		if resp.StatusCode != 200 {
			osexit.ExitBecauseError("No such command.")
		}
		defer resp.Body.Close()
		config, err := ioutil.ReadAll(resp.Body)
		osexit.ExitOnError(err)

		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)
		err = dkr.SaveConfig(home, command, config)
		osexit.ExitOnError(err)
		fmt.Printf("Command %s added.\n", color.GreenString(command))

		target, err := dkr.CopyProxy(command)
		osexit.ExitOnError(err)
		fmt.Printf("Proxy file for command %s created: %s\n",
			color.GreenString(command), color.GreenString(target))
	},
}
