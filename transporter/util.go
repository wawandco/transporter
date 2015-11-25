package transporter

import (
	"database/sql"
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
	return sql.Open("postgres", "user=transporter dbname=transporter sslmode=disable")
}

func MigrationIdentifier() int64 {
	return time.Now().UnixNano()
}
