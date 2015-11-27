package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/crufter/copyrecur"
)

type UpTemplateData struct {
	TempDir string
}

//Up runs any pending migration.
func Up(ctx *cli.Context) {
	temp := buildTempFolder()
	defer os.RemoveAll(temp)

	commandArgs := copyMigrationFiles(temp)
	main, _ := writeTemplateToFile(filepath.Join(temp, "main.go"), upTemplate, UpTemplateData{TempDir: temp})

	commandArgs = append(commandArgs, main)
	runTempFiles(commandArgs)
}

func runTempFiles(commandArgs []string) {
	cmd := exec.Command("go", commandArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if e := cmd.Run(); e != nil {
		log.Fatal("`go run` failed: ", e)
	}
}

func copyMigrationFiles(tempFolder string) []string {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	copyrecur.CopyFile(filepath.Join(base, "db", "config.yml"), filepath.Join(tempFolder, "config.yml"))
	commandArgs := []string{"run"}

	files, _ := ioutil.ReadDir(filepath.Join(base, "db", "migrations"))
	for _, file := range files {
		commandArgs = append(commandArgs, filepath.Join(tempFolder, file.Name()))
		copyrecur.CopyFile(filepath.Join(base, "db", "migrations", file.Name()), filepath.Join(tempFolder, file.Name()))
		replaceInFile(filepath.Join(tempFolder, file.Name()), "package migrations", "package main")
	}
	return commandArgs
}

func buildTempFolder() string {
	temp, e := ioutil.TempDir("", "transporter")
	if e != nil {
		log.Fatal(e)
	}
	return temp
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
	if err != nil {
		log.Println("Could not init database connection:", err)
	}

	transporter.RunAllMigrationsUp(db)
}
`))
