package cmd

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/wawandco/transporter/utils"
)

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
		TempDir:     temp,
		Environment: environment,
	}

	commandArgs := utils.CopyMigrationFilesTo(temp)
	main, _ := utils.WriteTemplateToFile(filepath.Join(temp, "main.go"), utils.UpTemplate, upTemplateData)

	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}
