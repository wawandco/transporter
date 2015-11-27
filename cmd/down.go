package cmd

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"
)

func Down(ctx *cli.Context) {
	temp := buildTempFolder()
	defer os.RemoveAll(temp)

	commandArgs := copyMigrationFiles(temp)
	main, _ := writeTemplateToFile(filepath.Join(temp, "main.go"), downTemplate, UpTemplateData{TempDir: temp})

	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}

var downTemplate = template.Must(template.New("down.template").Parse(`
package main

import (
	"log"
	"path/filepath"
	"github.com/wawandco/transporter/transporter"
	"io/ioutil"
)

func main() {
	log.Println("| Running Migrations Down")
	dat, _ := ioutil.ReadFile(filepath.Join("{{.TempDir}}","config.yml"))
	db, err := transporter.DBConnection(dat)
	if err != nil {
		log.Println("Could not init database connection:", err)
	}

	transporter.RunOneMigrationDown(db)
}
`))
