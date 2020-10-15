package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

const version = "0.4.0"

func init() {
	DkrCommand.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:                "version",
	Short:              "Displays version of dkr",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s\n", version)
	},
}
