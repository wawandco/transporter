package utils

import "text/template"

//MigrationTemplate is the template for the generation of the migrations.
const MigrationTemplate = `
package migrations
import (
  transporter "github.com/wawandco/transporter/core"
)

func init(){
  migration := transporter.Migration{
    Identifier: {{.Identifier}},
    Up: func(tx *transporter.Tx){
      tx.Exec("{{.UpCommand}}")
    },
    Down: func(tx *transporter.Tx){
      tx.Exec("{{.DownCommand}}")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Add(migration)
}
`

//UpTemplate is the template for the Main when cmd.Up runs.
var UpTemplate = template.Must(template.New("up.main.template").Parse(`
package main

import (
	"log"
	"path/filepath"
	transporter "github.com/wawandco/transporter/core"
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

//DownTemplate is the template for the main.go file when running Down cmd.
var DownTemplate = template.Must(template.New("down.main.template").Parse(`
package main

import (
	"log"
	"path/filepath"
	transporter "github.com/wawandco/transporter/core"
	"io/ioutil"
)

func main() {
	log.Println("| Running Migrations Down on [{{.Environment}}] environment")
	dat, _ := ioutil.ReadFile(filepath.Join("{{.TempDir}}","config.yml"))
	db, err := transporter.DBConnection(dat, "{{.Environment}}")

	if err != nil {
		log.Println("Could not init database connection:", err)
		return
	}

	defer db.Close()
	transporter.RunOneMigrationDown(db)
}
`))
