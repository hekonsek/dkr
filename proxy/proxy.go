package main

import (
	"github.com/hekonsek/dkr"
	"github.com/hekonsek/osexit"
	"os"
	"path"
)

func main() {
	binaryName := path.Base(os.Args[0])
	home, err := dkr.NewDkrHome()
	osexit.ExitOnError(err)
	config, err := dkr.ParseConfig(home, binaryName)
	osexit.ExitOnError(err)

	err = dkr.Sandbox(config.Image, config.Entrypoint, os.Args[1:]...)
	osexit.ExitOnError(err)
}
