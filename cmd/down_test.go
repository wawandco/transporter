package cmd

import (
	"flag"
	"log"
	"strconv"
	"testing"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	transporter "github.com/wawandco/transporter/core"
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

func TestDownBadMigration(t *testing.T) {
	for dname, man := range mans {
		utils.ClearTestTables(dname)
		utils.ClearTestMigrations()
		utils.SetupTestingFolders(dname)

		firstID := transporter.MigrationIdentifier()
		lastID := transporter.MigrationIdentifier()

		utils.GenerateMigrationFiles([]utils.TestMigration{
			utils.TestMigration{
				Identifier:  firstID,
				UpCommand:   "CREATE TABLE down_table (a varchar(255))",
				DownCommand: "DROP TABLE down_table",
			},
			utils.TestMigration{
				Identifier:  lastID,
				UpCommand:   "ALTER TABLE down_table ADD COLUMN o varchar(255)",
				DownCommand: "WUHdnns oius (Bad mig!);",
			},
		})

		context := cli.NewContext(nil, &flag.FlagSet{}, nil)
		Up(context)
		Down(context)

		con, _ := utils.BuildTestingConnection(dname)
		transporter.SetManager(man)

		version := transporter.DatabaseVersion(con)
		assert.Equal(t, version, strconv.FormatInt(lastID, 10))
		defer con.Close()
	}
}
