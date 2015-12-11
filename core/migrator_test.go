package core

import (
	"strconv"
	"testing"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/wawandco/transporter/managers"
	"github.com/wawandco/transporter/utils"
)

var sampleMigrations = []Migration{
	Migration{
		Identifier: MigrationIdentifier(),
		Up: func(tx *Tx) {
			tx.Exec("CREATE TABLE tests_table (a varchar(12))")
		},
		Down: func(tx *Tx) {
			tx.Exec("DROP TABLE tests_table;")
		},
	},

	Migration{
		Identifier: MigrationIdentifier(),
		Up: func(tx *Tx) {
			tx.Exec("ALTER table tests_table ADD COLUMN other varchar(20);")
		},
		Down: func(tx *Tx) {
			tx.Exec("ALTER table tests_table DROP COLUMN other;")
		},
	},
}

var mans = map[string]managers.DatabaseManager{
	"postgres": &managers.PostgreSQLManager{},
	"mysql":    &managers.MySQLManager{},
}

//TODO: multiple migrations with the same ID
func TestRegister(t *testing.T) {
	mig := Migration{
		Identifier: MigrationIdentifier(),
	}

	Add(mig)
	assert.Equal(t, 1, len(migrations))
}

func TestMigrationsTableDoesntExists(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)

		db, err := utils.BuildTestingConnection(name)
		defer db.Close()
		assert.Nil(t, err)
		assert.False(t, MigrationsTableExists(db))
	}
}

func TestMigrationsTableExists(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, err := utils.BuildTestingConnection(name)
		defer db.Close()
		assert.Nil(t, err)
		assert.True(t, MigrationsTableExists(db))
	}
}

func TestRunMigrationUp(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := sampleMigrations[0]

		RunMigrationUp(db, &m)
		_, err := db.Exec("Select * from tests_table;")
		assert.Nil(t, err)
		rows, _ := db.Query("Select * from " + MigrationsTable + ";")

		count := 0
		defer rows.Close()
		for rows.Next() {
			var identifier string
			rows.Scan(&identifier)
			id, _ := strconv.Atoi(identifier)
			assert.Equal(t, int64(id), m.Identifier)
			count++
		}

		assert.Equal(t, 1, count)
	}
}

func TestRunMigrationDown(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		m := sampleMigrations[0]

		RunMigrationUp(db, &m)
		RunMigrationDown(db, &m)

		_, err := db.Exec("Select * from tests_table;")
		assert.NotNil(t, err)

		rows, _ := db.Query("Select * from " + MigrationsTable + ";")
		count := 0
		defer rows.Close()
		for rows.Next() {
			var identifier string
			rows.Scan(&identifier)
			id, _ := strconv.Atoi(identifier)
			assert.Equal(t, id, m.Identifier)
			count++
		}

		assert.Equal(t, 0, count)
	}
}

func TestRunOneMigrationDown(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		migrations = []Migration{}

		m := sampleMigrations[0]
		Add(m)
		RunMigrationUp(db, &m)
		RunOneMigrationDown(db)

		_, err := db.Exec("Select * from tests_table;")
		assert.NotNil(t, err)

		rows, _ := db.Query("Select * from " + MigrationsTable + ";")
		count := 0
		defer rows.Close()
		for rows.Next() {
			count++
		}

		assert.Equal(t, 0, count)
	}
}

func TestRunAllMigrationsUp(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		Add(sampleMigrations[0])
		Add(sampleMigrations[1])
		RunAllMigrationsUp(db)

		_, err := db.Query("Select other from tests_table;")
		assert.Nil(t, err)

		rows, _ := db.Query("Select * from " + MigrationsTable + ";")
		count := 0
		defer rows.Close()
		for rows.Next() {
			count++
		}

		assert.Equal(t, count, 2)
	}
}

func TestRunAllMigrationsOnlyPending(t *testing.T) {
	for name, man := range mans {
		manager = man
		dropTables(name)
		createMigrationsTable(name)

		db, _ := utils.BuildTestingConnection(name)
		defer db.Close()

		migrations = []Migration{}

		Add(sampleMigrations[0])
		Add(sampleMigrations[1])

		db.Query("INSERT INTO " + MigrationsTable + " VALUES (" + sampleMigrations[1].GetID() + ");")
		RunAllMigrationsUp(db)

		_, err := db.Query("Select other from tests_table;")
		assert.NotNil(t, err)
	}

}
