package main

import (
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"github.com/spf13/cobra"
)

var cmdCreateCommandEntryPoint string

func init() {
	cmdCreateCommand.Flags().StringVarP(&cmdCreateCommandEntryPoint, "entrypoint", "", "", "")
	cmdParentCommand.AddCommand(cmdCreateCommand)
}

var cmdCreateCommand = &cobra.Command{
	Use: "add COMMAND IMAGE",
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		image := args[1]
		config := dkr.Config{Name: command, Image: image, Entrypoint: cmdCreateCommandEntryPoint}
		home, err := dkr.NewDkrHome()
		osexit.ExitOnError(err)
		err = config.Save(home)
		osexit.ExitOnError(err)
	},
}

