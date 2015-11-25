package transporter

import (
	"database/sql"
	"strconv"
)

//Migration represents migrations we will run up and down.
type Migration struct {
	Identifier int64
	Name       string
	Up         func(tx *sql.Tx)
	Down       func(tx *sql.Tx)
}

func (m *Migration) GetID() string {
	return strconv.FormatInt(m.Identifier, 10)
}

func (m *Migration) Pending(db *sql.DB) bool {
	rows, _ := db.Query("Select * from " + MigrationsTable + " WHERE identifier =" + m.GetID())
	defer rows.Close()
	count := 0

	for rows.Next() {
		count++
	}

	return count == 0
}

type ByIdentifier []Migration

func (a ByIdentifier) Len() int           { return len(a) }
func (a ByIdentifier) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIdentifier) Less(i, j int) bool { return a[i].Identifier < a[j].Identifier }
