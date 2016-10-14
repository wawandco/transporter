package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/wawandco/transporter/cmd"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "database, d",
			Usage: "Database URL",
			Destination: &cmd.DatabaseURL,
		},

		cli.StringFlag{
			Name:  "driver, i",
			Usage: "Database driver name",
			Destination: &cmd.DriverName,
		},
	}

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
		{
			Name:    "up",
			Aliases: []string{},
			Usage:   "transporter up",
			Action:  cmd.Up,
			Flags
		},
		{
			Name:    "down",
			Aliases: []string{},
			Usage:   "transporter down",
			Action:  cmd.Down,
		},
	}

	app.Run(os.Args)
}
