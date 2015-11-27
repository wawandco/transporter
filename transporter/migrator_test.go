package transporter

import (
	"database/sql"
	"strconv"
	"testing"

	_ "github.com/wawandco/transporter/Godeps/_workspace/src/github.com/lib/pq"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

var sampleMigrations = []Migration{
	Migration{
		Identifier: MigrationIdentifier(),
		Up: func(tx *sql.Tx) {
			tx.Exec("Create table tests_table (a varchar);")
		},
		Down: func(tx *sql.Tx) {
			tx.Exec("Drop table tests_table;")
		},
	},

	Migration{
		Identifier: MigrationIdentifier(),
		Up: func(tx *sql.Tx) {
			tx.Exec("ALTER table tests_table ADD COLUMN other varchar(20);")
		},
		Down: func(tx *sql.Tx) {
			tx.Exec("ALTER table tests_table DROP COLUMN other;")
		},
	},
}

//TODO: multiple migrations with the same ID
func TestRegister(t *testing.T) {
	mig := Migration{
		Identifier: MigrationIdentifier(),
	}

	Register(mig)
	assert.Equal(t, 1, len(migrations))
}

func TestMigrationsTableDoesntExists(t *testing.T) {
	dropTables()
	db, err := testConnection()
	defer db.Close()
	assert.Nil(t, err)
	assert.False(t, MigrationsTableExists(db))
}

func TestMigrationsTableExists(t *testing.T) {
	dropTables()
	createMigrationsTable()
	db, err := testConnection()
	defer db.Close()

	assert.Nil(t, err)
	assert.True(t, MigrationsTableExists(db))
}

func TestRunMigrationUp(t *testing.T) {
	dropTables()
	createMigrationsTable()
	db, _ := testConnection()
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

func TestRunMigrationDown(t *testing.T) {
	dropTables()
	createMigrationsTable()
	db, _ := testConnection()
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

func TestRunOneMigrationDown(t *testing.T) {
	dropTables()
	createMigrationsTable()
	db, _ := testConnection()
	defer db.Close()
	migrations = []Migration{}

	m := sampleMigrations[0]
	Register(m)
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

func TestRunAllMigrationsUp(t *testing.T) {
	dropTables()
	createMigrationsTable()
	migrations = []Migration{}

	Register(sampleMigrations[0])
	Register(sampleMigrations[1])

	db, _ := testConnection()
	RunAllMigrationsUp(db)

	defer db.Close()
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

func TestRunAllMigrationsOnlyPending(t *testing.T) {
	dropTables()
	createMigrationsTable()
	migrations = []Migration{}

	db, _ := testConnection()
	defer db.Close()

	Register(sampleMigrations[0])
	Register(sampleMigrations[1])

	db.Query("INSERT INTO " + MigrationsTable + " VALUES (" + sampleMigrations[1].GetID() + ");")
	RunAllMigrationsUp(db)

	_, err := db.Query("Select other from tests_table;")
	assert.NotNil(t, err)
}
