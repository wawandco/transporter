package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/managers"
	"github.com/wawandco/transporter/utils"
)

func TestDropMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.Exec("CREATE TABLE tests_table (a varchar(12))")
			},
			Down: func(tx *Tx) {
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)
		RunMigrationDown(db, &m)
		_, err := db.Exec("SELECT * FROM tests_table")
		assert.NotNil(t, err)

	}
}

func TestCreateTableMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.CreateTable("tests_table", managers.Table{
					"a":            "varchar(12)",
					"other_column": "integer",
					"float_column": "float",
				})
			},
			Down: func(tx *Tx) {
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)
		_, err := db.Exec("SELECT * FROM tests_table")
		assert.Nil(t, err)
	}
}

func TestAddColumnToTableMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.CreateTable("tests_table", managers.Table{
					"a":            "varchar(12)",
					"other_column": "integer",
					"float_column": "float",
				})

				tx.AddColumn("tests_table", "lately_added", "varchar(255)")
			},
			Down: func(tx *Tx) {
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)
		_, err := db.Exec("SELECT lately_added FROM tests_table")
		assert.Nil(t, err)
	}
}

func TestDropColumnMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.CreateTable("tests_table", managers.Table{
					"a":            "varchar(12)",
					"other_column": "integer",
					"float_column": "float",
				})

				tx.DropColumn("tests_table", "float_column")
			},
			Down: func(tx *Tx) {
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)
		_, err := db.Exec("SELECT float_column FROM tests_table")
		assert.NotNil(t, err)
	}
}

func TestChangeColumnMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.CreateTable("tests_table", managers.Table{
					"a":            "varchar(12)",
					"other_column": "integer",
					"float_column": "float",
				})

				tx.ChangeColumnType("tests_table", "float_column", "integer")
			},
			Down: func(tx *Tx) {
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)
		_, err := db.Exec("SELECT float_column FROM tests_table")
		assert.Nil(t, err)

		//TODO: need to test this properly.
	}
}

func TestChangeColumnNameMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.CreateTable("tests_table", managers.Table{
					"a":            "varchar(12)",
					"other_column": "integer",
					"float_column": "float",
				})

				tx.RenameColumn("tests_table", "float_column", "point_column")
			},
			Down: func(tx *Tx) {
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)
		_, err := db.Exec("SELECT point_column FROM tests_table")
		if name == "mysql" {
			assert.NotNil(t, err) //Note: Mysql requires the column type.
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestRenameTableMigration(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := Migration{
			Identifier: MigrationIdentifier(),
			Up: func(tx *Tx) {
				tx.CreateTable("tests_table", managers.Table{
					"a":            "varchar(12)",
					"other_column": "integer",
					"float_column": "float",
				})

				tx.RenameTable("tests_table", "pthkkk_table")
			},
			Down: func(tx *Tx) {
				tx.RenameTable("pthkkk_table", "tests_table")
				tx.DropTable("tests_table")
			},
		}

		RunMigrationUp(db, &m)

		if name != "mysql" {
			_, err := db.Exec("SELECT other_column FROM tests_table")
			assert.NotNil(t, err)

			_, err = db.Exec("SELECT other_column FROM pthkkk_table")
			assert.Nil(t, err)
		}

	}
}
