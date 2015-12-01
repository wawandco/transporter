package cmd

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/utils"
)

//Up runs any pending migration.
func Up(ctx *cli.Context) {
	temp := buildTempFolder()
	defer os.RemoveAll(temp)

	var environment string
	if len(ctx.Args()) == 0 || ctx.Args()[0] == "" {
		environment = "development"
	} else {
		environment = ctx.Args()[0]
	}

	upTemplateData := CmdTemplateData{
		TempDir:     temp,
		Environment: environment, //TODO: this should come from ctx or be development
	}

	commandArgs := utils.CopyMigrationFilesTo(temp)
	main, _ := utils.WriteTemplateToFile(filepath.Join(temp, "main.go"), upTemplate, upTemplateData)

	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}

var upTemplate = template.Must(template.New("main.template").Parse(`
package main

import (
	"log"
	"path/filepath"
	"github.com/wawandco/transporter/transporter"
	"io/ioutil"
)

func main() {
	log.Println("| Running Migrations UP on [{{.Environment}}] environment")
	dat, _ := ioutil.ReadFile(filepath.Join("{{.TempDir}}","config.yml"))
	db, err := transporter.DBConnection(dat, "{{.Environment}}")

	if err != nil {
		log.Println("Could not init database connection:", err)
		return
	}

	defer db.Close()
	transporter.RunAllMigrationsUp(db)
}
`))
