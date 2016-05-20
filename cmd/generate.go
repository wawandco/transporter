package cmd

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"strconv"

	"github.com/codegangsta/cli"
	"github.com/serenize/snaker"
	"github.com/wawandco/transporter/utils"
)

//MigrationData is used to store migration data sent to the template.
type MigrationData struct {
	Identifier  int64
	UpCommand   string
	DownCommand string
	Name        string
}

//Generate generates a migration on the migrations folder
func Generate(c *cli.Context) {
	setupTestingEnv()
	base := os.Getenv("TRANS_TESTING_FOLDER")

	name := "migration"
	if c.App != nil && len(c.Args()) > 0 {
		name = c.Args().First()
	}

	name = snaker.CamelToSnake(name)
	identifier := time.Now().UnixNano()

	migration := MigrationData{
		Identifier: identifier,
		Name:       name,
	}

	buff := bytes.NewBufferString("")
	tmpl, _ := template.New("migration").Parse(utils.MigrationTemplate)
	_ = tmpl.Execute(buff, migration)

	fileName := strconv.FormatInt(identifier, 10) + "_" + name + ".go"
	path := filepath.Join(base, "db", "migrations", fileName)

	err := ioutil.WriteFile(path, buff.Bytes(), generatedFilePermissions)
	if err != nil {
		log.Println(err)
		log.Println("| Could not write migration file, please check db/migrations folder exists")
	}
}
