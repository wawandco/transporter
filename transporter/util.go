package transporter

import (
	"database/sql"
	"os"
	"time"
)

func dropTables() {
	db, _ := testConnection()
	db.Exec("DROP TABLE IF EXISTS  " + MigrationsTable + ";")
	db.Exec("DROP TABLE IF EXISTS tests_table ;")
}

func createMigrationsTable() {
	db, _ := testConnection()
	db.Exec("CREATE TABLE IF NOT EXISTS  " + MigrationsTable + " ( identifier decimal NOT NULL );")
}

func CreateMigrationsTable(db *sql.DB) {
	db.Exec("CREATE TABLE IF NOT EXISTS  " + MigrationsTable + " ( identifier decimal NOT NULL );")
}

func testConnection() (*sql.DB, error) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		url = "user=transporter dbname=transporter sslmode=disable"
	}
	return sql.Open("postgres", url)
}

func MigrationIdentifier() int64 {
	return time.Now().UnixNano()
}
