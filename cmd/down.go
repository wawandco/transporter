package cmd

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/wawandco/transporter/utils"
)

//Down is the command runt when you do `tranporter down` on your CLI.
func Down(ctx *cli.Context) {
	if Count == -1 {
		Count = 1 //Default value for Down, one migration
	}

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
		Count:       Count,
	}

	commandArgs := utils.CopyMigrationFilesTo(temp)
	main, _ := utils.WriteTemplateToFile(filepath.Join(temp, "main.go"), utils.DownTemplate, downTemplateData)

	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}
