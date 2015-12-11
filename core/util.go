package core

import (
	"database/sql"
	"time"

	"github.com/wawandco/transporter/utils"
)

func dropTables(driver string) {
	db, _ := utils.BuildTestingConnection(driver)
	db.Exec("DROP TABLE IF EXISTS  " + MigrationsTable + ";")
	db.Exec("DROP TABLE IF EXISTS tests_table ;")
}

func createMigrationsTable(driver string) {
	db, _ := utils.BuildTestingConnection(driver)
	query := manager.CreateMigrationsTableQuery(MigrationsTable)
	db.Exec(query)
}

// CreateMigrationsTable creates the migrations table based on the driver specific sql.
func CreateMigrationsTable(db *sql.DB) {
	query := manager.CreateMigrationsTableQuery(MigrationsTable)
	db.Exec(query)
}

//MigrationIdentifier returns a unixnano used to identify the order of the migration.
func MigrationIdentifier() int64 {
	return time.Now().UnixNano()
}

func driverRegistered(e string) bool {
	result := false

	for _, a := range sql.Drivers() {
		if a == e {
			result = true
			break
		}
	}

	return result
}
