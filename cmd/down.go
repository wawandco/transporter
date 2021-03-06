package cmd

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"github.com/wawandco/transporter/utils"
)

//Down is the command runt when you do `tranporter down` on your CLI.
func Down(ctx *cli.Context) {
	temp := buildTempFolder()
	defer os.RemoveAll(temp)

	var environment string
	if len(ctx.Args()) == 0 || ctx.Args()[0] == "" {
		environment = "development"
	} else {
		environment = ctx.Args()[0]
	}

	downTemplateData := MainData{
		TempDir:     temp,
		Environment: environment,
	}

	commandArgs := utils.CopyMigrationFilesTo(temp)
	template := utils.DownTemplate

	if DatabaseURL != "" {
		template = utils.DownFlagsTemplate
	}

	main, _ := utils.WriteTemplateToFile(filepath.Join(temp, "main.go"), template, downTemplateData)

	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}
