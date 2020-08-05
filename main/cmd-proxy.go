package main

import (
	"fmt"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
	"io/ioutil"
	"github.com/fatih/color"
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

		proxyBytes, err := ioutil.ReadFile("/usr/bin/dkr-proxy")
		osexit.ExitOnError(err)
		target := "/usr/bin/" + command
		err = ioutil.WriteFile(target, proxyBytes, 0555)
		osexit.ExitOnError(err)

		fmt.Printf("Proxy file for command %s created: %s\n",
			color.GreenString(command) , color.GreenString(target))
	},
}
