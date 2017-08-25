package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"os"
)

func SimpleCli() {
	app := cli.NewApp()
	app.Run(os.Args)
}