package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"text/template"

	"gopkg.in/yaml.v2"

	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/transporter"
)

const migTPL = `
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: {{.Identifier}},
    Up: func(tx *sql.Tx){
      tx.Exec("{{.UpCommand}}")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("{{.DownCommand}}")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
`

type TestMig struct {
	Identifier  int64
	UpCommand   string
	DownCommand string
}

func TestUp(t *testing.T) {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	defer os.RemoveAll(filepath.Join(base, "db", "migrations"))

	setupTestingEnv()
	cleanTables()

	buildBaseFolders()
	buildTestConfig()
	generateTestMigrations([]TestMig{
		TestMig{
			Identifier: transporter.MigrationIdentifier(),
			UpCommand:  "Create table other_table (a varchar(255) );",
		},
		TestMig{
			Identifier: transporter.MigrationIdentifier(),
			UpCommand:  "Alter table other_table add column o varchar(12);",
		},
	})

	context := cli.Context{}
	Up(&context)

	con, _ := buildConnectionFromConfig()
	_, err := con.Query("Select * from other_table;")
	assert.Nil(t, err)
}

func TestUpBadMigration(t *testing.T) {
	setupTestingEnv()
	cleanTables()

	buildBaseFolders()
	buildTestConfig()
	generateTestMigrations([]TestMig{
		TestMig{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Create table other_table (a varchar(255) );",
			DownCommand: "Drop Table other_table;",
		},
		TestMig{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Alter table shshshshs other_table add column o varchar(12);",
			DownCommand: "",
		},
		TestMig{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Alter table other_table drop column a;",
			DownCommand: "",
		},
	})

	context := cli.Context{}
	Up(&context)

	con, _ := buildConnectionFromConfig()
	_, err := con.Query("Select a from other_table;")
	assert.Nil(t, err)
}

func buildBaseFolders() {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.Mkdir(filepath.Join(base), generatedFilePermissions)
	os.Mkdir(filepath.Join(base, "db"), generatedFilePermissions)
	os.Mkdir(filepath.Join(base, "db", "migrations"), generatedFilePermissions)
}

func generateTestMigrations(migs []TestMig) {
	base := os.Getenv("TRANS_TESTING_FOLDER")

	for _, mig := range migs {
		buff := bytes.NewBufferString("")
		tmpl, _ := template.New("test_migration").Parse(migTPL)
		_ = tmpl.Execute(buff, mig)

		fileName := strconv.FormatInt(mig.Identifier, 10) + "_a.go"
		path := filepath.Join(base, "db", "migrations", fileName)
		ioutil.WriteFile(path, buff.Bytes(), generatedFilePermissions)
	}
}

func buildTestConfig() {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	config := Config{
		Database: map[string]string{
			"url":    "user=transporter dbname=transporter sslmode=disable",
			"driver": "postgres",
		},
	}

	d, _ := yaml.Marshal(&config)
	path := filepath.Join(base, "db", "config.yml")
	ioutil.WriteFile(path, d, generatedFilePermissions)
}
