package transporter

import (
	"database/sql"
	"os"
	"strings"
	"time"
)

func dropTables(driver string) {
	db, _ := testConnection(driver)
	db.Exec("DROP TABLE IF EXISTS  " + MigrationsTable + ";")
	db.Exec("DROP TABLE IF EXISTS tests_table ;")
}

func createMigrationsTable(driver string) {
	db, _ := testConnection(driver)
	query := manager.CreateMigrationsTableQuery(MigrationsTable)
	db.Exec(query)
}

func CreateMigrationsTable(db *sql.DB) {
	query := manager.CreateMigrationsTableQuery(MigrationsTable)
	db.Exec(query)
}

func testConnection(driver string) (*sql.DB, error) {
	url := os.Getenv(strings.ToUpper(driver) + "_DATABASE_URL")
	return sql.Open(driver, url)
}

func MigrationIdentifier() int64 {
	return time.Now().UnixNano()
}
