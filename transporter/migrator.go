package transporter

import (
	"database/sql"
	"log"
	"sort"
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

// //TODO: Implement me
// func MigrationsComplete(db *sql.DB) bool {
// 	return false
// }
//
//TODO: Implement me

func RunAllMigrationsUp(db *sql.DB) {

	//1. Migrations table exists? -> Create if needed
	if !MigrationsTableExists(db) {
		CreateMigrationsTable(db)
	}

	sort.Sort(ByIdentifier(migrations))
	for _, migration := range migrations {
		// Check if migration is on the database already
		if migration.Pending(db) {
			RunMigrationUp(db, &migration)
		}
	}
}

// RunMigrationUp runs a single migration up and if success it saves the
// Migration identifier on the migrations table.
func RunMigrationUp(db *sql.DB, m *Migration) {
	tx, err := dbTransaction(db)
	if err != nil {
		return
	}

	if m.Up != nil {
		m.Up(tx)
		err = tx.Commit()

		if err == nil {
			_, err = db.Exec("INSERT INTO " + MigrationsTable + "( identifier ) VALUES ('" + m.GetID() + "') ;")
		} else {
			log.Println("| Error, Could not complete your migration (" + m.GetID() + "), please check your sql.")
			//TODO: Inform the SQL error
		}
	}
}

// RunMigrationDown Runs one migration down and if successful it removes the migration from
// the migrations table.
func RunMigrationDown(db *sql.DB, m *Migration) {
	tx, err := dbTransaction(db)
	if err != nil {
		return
	}

	m.Down(tx)
	err = tx.Commit()

	if err == nil {
		_, err = db.Exec("DELETE FROM " + MigrationsTable + " WHERE identifier = '" + m.GetID() + "' ;")
	} else {
		log.Println("| Error, Could not complete your migration (" + m.GetID() + "), please check your sql.")
		//TODO: Inform the SQL error
	}
}

func dbTransaction(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("| Error, could not initialize transaction")
	}

	return tx, err
}

//TODO: Implement me
func DownAllMigrations(db *sql.DB) {}
