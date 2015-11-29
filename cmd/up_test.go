package cmd

import (
	"testing"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/transporter"
	"github.com/wawandco/transporter/utils"
)

func TestUp(t *testing.T) {
	utils.ClearTestTables()
	utils.ClearTestMigrations()
	utils.SetupTestingFolders()

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

	context := cli.Context{}
	Up(&context)

	con, _ := utils.BuildTestingConnection()
	_, err := con.Query("Select * from other_table;")
	assert.Nil(t, err)
}

func TestUpBadMigration(t *testing.T) {
	utils.ClearTestTables()
	utils.ClearTestMigrations()
	utils.SetupTestingFolders()

	utils.GenerateMigrationFiles([]utils.TestMigration{
		utils.TestMigration{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Create table other_table (a varchar(255) );",
			DownCommand: "Drop Table other_table;",
		},
		utils.TestMigration{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Alter table shshshshs other_table add column o varchar(12);",
			DownCommand: "",
		},
		utils.TestMigration{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Alter table other_table drop column a;",
			DownCommand: "",
		},
	})

	context := cli.Context{}
	Up(&context)

	con, _ := utils.BuildTestingConnection()
	_, err := con.Query("Select a from other_table;")
	assert.Nil(t, err)
}
