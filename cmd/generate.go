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

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/serenize/snaker"
	"github.com/wawandco/transporter/transporter"
)

var migrationTPL = `
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: {{.Identifier}},
    Up: func(tx *sql.Tx){
      //you can use here tx.Exec to change your DB up
    },
    Down: func(tx *sql.Tx){
      //you can use here tx.Exec to change your DB down
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
`

//Generate generates a migration on the migrations folder
func Generate(c *cli.Context) {
	base := ""
	if os.Getenv("TRANS_TESTING_FOLDER") != "" {
		base = os.Getenv("TRANS_TESTING_FOLDER")
		os.Mkdir(base, generatedFilePermissions)
	}

	name := "migration"
	if c.App != nil && len(c.Args()) > 0 {
		name = c.Args().First()
	}

	name = snaker.CamelToSnake(name)
	identifier := time.Now().UnixNano()

	migration := transporter.Migration{
		Identifier: identifier,
		Name:       name,
	}

	buff := bytes.NewBufferString("")
	tmpl, _ := template.New("migration").Parse(migrationTPL)
	_ = tmpl.Execute(buff, migration)

	fileName := strconv.FormatInt(identifier, 10) + "_" + name + ".go"
	path := filepath.Join(base, "db", "migrations", fileName)

	err := ioutil.WriteFile(path, buff.Bytes(), generatedFilePermissions)
	if err != nil {
		log.Println("| Could not write migration file, please check db/migrations folder exists")
	}
}
