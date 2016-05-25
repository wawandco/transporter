package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/wawandco/transporter/cmd"
)

var version = "0.0.1"

func main() {

	app := cli.NewApp()

	app.Name = "transporter"
	app.Version = version

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
		{
			Name:    "up",
			Aliases: []string{},
			Usage:   "transporter up",
			Action:  cmd.Up,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "count",
					Destination: &cmd.Count,
				},
			},
		},
		{
			Name:    "down",
			Aliases: []string{},
			Usage:   "transporter down",
			Action:  cmd.Down,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "count",
					Destination: &cmd.Count,
				},
			},
		},
	}

	app.Run(os.Args)
}
