package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/wawandco/transporter/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "transporter"
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"ini"},
			Usage:   "transporter init",
			Action:  cmd.Init,
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "transporter generate [NAME]",
			Action:  cmd.Generate,
		},
	}

	app.Run(os.Args)
}
