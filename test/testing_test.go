package test

import (
	"database/sql"
	"os"
)

func buildTestingConnection() (*sql.DB, error) {
	url := os.Getenv("TEST_DATABASE_URL")
	return sql.Open("postgres", url)
}

var testingTables = []string{
	"other_table",
	"down_table",
	"transporter_migrations",
}

func ClearTestTables() {
	conn, _ := buildTestingConnection()
	defer conn.Close()

	for _, t := range testingTables {
		conn.Exec("DROP TABLE IF EXISTS  " + t + ";")
	}
}
