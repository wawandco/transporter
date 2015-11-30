package cmd

import (
	"log"
	"testing"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/transporter"
	"github.com/wawandco/transporter/utils"
)

func TestDown(t *testing.T) {
	utils.ClearTestTables()
	utils.ClearTestMigrations()
	utils.SetupTestingFolders()

	utils.GenerateMigrationFiles([]utils.TestMigration{
		utils.TestMigration{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Create table down_table (a varchar(255) );",
			DownCommand: "Drop table down_table;",
		},
		utils.TestMigration{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Alter table down_table add column o varchar(12);",
			DownCommand: "Alter table down_table drop column o;",
		},
	})

	context := cli.Context{}
	Up(&context)
	Down(&context)

	con, _ := utils.BuildTestingConnection()
	defer con.Close()

	_, err := con.Query("Select a from down_table;")
	assert.Nil(t, err)

	Down(&context)

	_, err = con.Query("Select a from down_table;")
	log.Println(err)
	assert.NotNil(t, err)

}
