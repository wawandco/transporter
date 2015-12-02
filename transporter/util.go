package transporter

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

func CreateMigrationsTable(db *sql.DB) {
	query := manager.CreateMigrationsTableQuery(MigrationsTable)
	db.Exec(query)
}

func MigrationIdentifier() int64 {
	return time.Now().UnixNano()
}
