package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/transporter"
)

func TestDown(t *testing.T) {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	defer os.RemoveAll(filepath.Join(base, "db", "migrations"))

	setupTestingEnv()
	cleanTables()

	buildBaseFolders()
	buildTestConfig()
	generateTestMigrations([]TestMig{
		TestMig{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Create table down_table (a varchar(255) );",
			DownCommand: "Drop table down_table;",
		},
		TestMig{
			Identifier:  transporter.MigrationIdentifier(),
			UpCommand:   "Alter table down_table add column o varchar(12);",
			DownCommand: "Alter table down_table drop column o;",
		},
	})

	context := cli.Context{}
	Up(&context)
	Down(&context)

	con, _ := buildConnectionFromConfig()
	_, err := con.Query("Select a from down_table;")
	assert.Nil(t, err)

	_, err = con.Query("Select o from down_table;")
	assert.NotNil(t, err)

	Down(&context)

	_, err = con.Query("Select a from down_table;")
	assert.NotNil(t, err)

}
