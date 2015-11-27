package transporter

import (
	"database/sql"
	"errors"
	"log"
	"sort"
	"strconv"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v1"
)

var migrations []Migration

const (
	// MigrationsTable is the Db table where we will store migrations
	MigrationsTable = "transporter_migrations"
)

// Register function adds a migration to the migrations array
// So Transporter verifies if the migration is already completed or runs
// the desired migration as needed.
//
// Register checks if migration attempted to be registered id is unique.
func Register(m Migration) {
	migrations = append(migrations, m)
}

// MigrationsTableExists returns true if the table for the migrations already exists.
func MigrationsTableExists(db *sql.DB) bool {
	_, err := db.Query("SELECT * FROM " + MigrationsTable)
	return err == nil
}

//RunAllMigrationsUp runs pending migrations and stores migration on the migrations table.
func RunAllMigrationsUp(db *sql.DB) {

	//1. Migrations table exists? -> Create if needed
	if !MigrationsTableExists(db) {
		CreateMigrationsTable(db)
	}

	sort.Sort(ByIdentifier(migrations))
	var err error
	for _, migration := range migrations {
		// Check if migration is on the database already
		if migration.Pending(db) {
			err = RunMigrationUp(db, &migration)

			if err != nil {
				log.Println("| Error: " + err.Error())
				break
			}
		}
	}

	version := DatabaseVersion(db)
	if err == nil && version != "" {
		log.Println("| Done, new database version is " + version)
	}
}

// RunMigrationUp runs a single migration up and if success it saves the
// Migration identifier on the migrations table.
func RunMigrationUp(db *sql.DB, m *Migration) error {
	tx, err := dbTransaction(db)
	if err != nil {
		return errors.New("Could not open a db transaction.")
	}

	if m.Up != nil {
		m.Up(tx)
		err = tx.Commit()

		if err != nil {
			log.Println("| Error, Could not complete your migration (" + m.GetID() + "), please check your sql.")
			return errors.New("Could not complete your migration (" + m.GetID() + "), please check your sql.")
		}

		_, err = db.Exec("INSERT INTO " + MigrationsTable + "( identifier ) VALUES ('" + m.GetID() + "') ;")
		return nil
	}

	return errors.New("Migration doesnt have Up function defined.")
}

// RunMigrationDown Runs one migration down and if successful it removes the migration from
// the migrations table.
func RunMigrationDown(db *sql.DB, m *Migration) error {
	tx, err := dbTransaction(db)
	if err != nil {
		return errors.New("Could not begin a transaction.")
	}

	if m.Down != nil {
		m.Down(tx)
		err = tx.Commit()

		if err != nil {
			log.Println("| Error, Could not complete your migration (" + m.GetID() + "), please check your sql.")
			return errors.New("Could not complete your migration (" + m.GetID() + "), please check your sql.")
		}

		_, err = db.Exec("DELETE FROM " + MigrationsTable + " WHERE identifier = '" + m.GetID() + "' ;")
		return nil
	}

	return errors.New("Migration (" + m.GetID() + ") doesn't have Down function defined.")
}

//DownOneMigration Run down last migration that have completed.
func RunOneMigrationDown(db *sql.DB) {
	if !MigrationsTableExists(db) {
		CreateMigrationsTable(db)
	}

	sort.Sort(ByIdentifier(migrations))
	identifier := DatabaseVersion(db)
	id, err := strconv.Atoi(identifier)

	if err != nil {
		log.Println("Sorry, there is no migration to run back")
	}

	for _, mig := range migrations {
		if mig.Identifier-int64(id) == 0 {
			log.Println("| Running " + mig.GetID() + " back")
			err := RunMigrationDown(db, &mig)
			if err == nil {
				version := DatabaseVersion(db)
				if version != "" {
					log.Println("| Done, new database version is " + version)
				} else {
					log.Println("| Done, All existing migrations down.")
				}
			}
			break
		}
	}

}

func DatabaseVersion(db *sql.DB) string {
	rows, _ := db.Query("Select max(identifier) from " + MigrationsTable + ";")
	var identifier string
	for rows.Next() {
		rows.Scan(&identifier)
		break
	}

	return identifier
}

//DBConnection Returns a DB connection from the yml config file
func DBConnection(ymlFile []byte) (*sql.DB, error) {
	type Config struct {
		Database map[string]string
	}

	var config Config
	err := yaml.Unmarshal([]byte(ymlFile), &config)

	if err != nil {
		return nil, err
	}

	return sql.Open(config.Database["driver"], config.Database["url"])
}

func dbTransaction(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("| Error, could not initialize transaction")
	}

	return tx, err
}
