package cmd

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"

	transporter "github.com/wawandco/transporter/core"
	"github.com/wawandco/transporter/utils"
)

func TestUp(t *testing.T) {
	for dname := range mans {
		utils.ClearTestTables(dname)
		utils.ClearTestMigrations()
		utils.SetupTestingFolders(dname)

		utils.GenerateMigrationFiles([]utils.TestMigration{
			utils.TestMigration{
				Identifier: transporter.MigrationIdentifier(),
				UpCommand:  "Create table other_table (a varchar(255) );",
			},
			utils.TestMigration{
				Identifier: transporter.MigrationIdentifier(),
				UpCommand:  "Alter table other_table add column o varchar(12);",
			},
		})

		context := cli.NewContext(nil, &flag.FlagSet{}, nil)
		Up(context)

		con, _ := utils.BuildTestingConnection(dname)
		_, err := con.Query("Select * from other_table;")
		assert.Nil(t, err)
	}
}

func TestUpBadMigration(t *testing.T) {
	for dname := range mans {
		utils.ClearTestTables(dname)
		utils.ClearTestMigrations()
		utils.SetupTestingFolders(dname)

		utils.GenerateMigrationFiles([]utils.TestMigration{
			utils.TestMigration{
				Identifier:  transporter.MigrationIdentifier(),
				UpCommand:   "Create table other_table (a varchar(255) );",
				DownCommand: "Drop Table other_table;",
			},
			utils.TestMigration{
				Identifier:  transporter.MigrationIdentifier(),
				UpCommand:   "Alter table shshshshs other_table add column o varchar(12);",
				DownCommand: "Drop column Wihii;",
			},
			utils.TestMigration{
				Identifier:  transporter.MigrationIdentifier(),
				UpCommand:   "Alter table other_table drop column a;",
				DownCommand: "",
			},
		})

		context := cli.NewContext(nil, &flag.FlagSet{}, nil)
		Up(context)

		con, _ := utils.BuildTestingConnection(dname)
		_, err := con.Query("Select a from other_table;")
		assert.Nil(t, err)
	}
}
