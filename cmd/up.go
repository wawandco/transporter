package cmd

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"github.com/wawandco/transporter/utils"
)

var DatabaseURL string
var DriverName string

//Up is the command runt when you do `tranporter up` on your CLI.
func Up(ctx *cli.Context) {
	temp := buildTempFolder()
	defer os.RemoveAll(temp)

	var environment string
	if len(ctx.Args()) == 0 || ctx.Args()[0] == "" {
		environment = "development"
	} else {
		environment = ctx.Args()[0]
	}

	upTemplateData := MainData{
		TempDir:        temp,
		Environment:    environment,
		DatabaseURL:    DatabaseURL,
		DatabaseDriver: DriverName,
	}

	template := utils.UpTemplate
	commandArgs := utils.CopyMigrationFilesTo(temp)

	if DatabaseURL != "" {
		template = utils.UpFlagsTemplate
	}

	main, _ := utils.WriteTemplateToFile(filepath.Join(temp, "main.go"), template, upTemplateData)
	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}
