package cmd

import (
	"flag"
	"log"
	"testing"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/transporter"
	"github.com/wawandco/transporter/utils"
)

func TestDown(t *testing.T) {
	for dname := range mans {
		utils.ClearTestTables(dname)
		utils.ClearTestMigrations()
		utils.SetupTestingFolders(dname)

		utils.GenerateMigrationFiles([]utils.TestMigration{
			utils.TestMigration{
				Identifier:  transporter.MigrationIdentifier(),
				UpCommand:   "CREATE TABLE down_table (a varchar(255))",
				DownCommand: "DROP TABLE down_table",
			},
			utils.TestMigration{
				Identifier:  transporter.MigrationIdentifier(),
				UpCommand:   "ALTER TABLE down_table ADD COLUMN o varchar(255)",
				DownCommand: "ALTER TABLE down_table DROP COLUMN o",
			},
		})

		context := cli.NewContext(nil, &flag.FlagSet{}, nil)
		Up(context)
		Down(context)

		con, _ := utils.BuildTestingConnection(dname)
		defer con.Close()

		_, err := con.Query("Select a from down_table")
		assert.Nil(t, err)

		Down(context)

		_, err = con.Query("Select a from down_table")
		log.Println(err)
		assert.NotNil(t, err)
	}
}
