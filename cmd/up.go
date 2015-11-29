package cmd

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/utils"
)

type UpTemplateData struct {
	TempDir string
}

//Up runs any pending migration.
func Up(ctx *cli.Context) {
	temp := buildTempFolder()
	defer os.RemoveAll(temp)

	commandArgs := utils.CopyMigrationFilesTo(temp)
	main, _ := utils.WriteTemplateToFile(filepath.Join(temp, "main.go"), upTemplate, UpTemplateData{TempDir: temp})

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
	log.Println("| Running Migrations UP")
	dat, _ := ioutil.ReadFile(filepath.Join("{{.TempDir}}","config.yml"))
	db, err := transporter.DBConnection(dat)
	defer db.Close()

	if err != nil {
		log.Println("Could not init database connection:", err)
	}

	transporter.RunAllMigrationsUp(db)
}
`))
