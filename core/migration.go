package core

import (
	"database/sql"
	"strconv"
)

//Migration represents migrations we will run up and down.
type Migration struct {
	Identifier int64
	Name       string
	Up         func(tx *Tx)
	Down       func(tx *Tx)
}

//GetID Returns a string representation of the migration identifier.
func (m *Migration) GetID() string {
	return strconv.FormatInt(m.Identifier, 10)
}

//Pending returns if a particular migration is pending.
// TODO: move it to be vendor specific inside the managers.
func (m *Migration) Pending(db *sql.DB) bool {
	rows, _ := db.Query("Select * from " + MigrationsTable + " WHERE identifier =" + m.GetID())
	defer rows.Close()
	count := 0

	for rows.Next() {
		count++
	}

	return count == 0
}

//ByIdentifier is a specific order that causes first-created migrations to run first based on the identifier.
type ByIdentifier []Migration

func (a ByIdentifier) Len() int           { return len(a) }
func (a ByIdentifier) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIdentifier) Less(i, j int) bool { return a[i].Identifier < a[j].Identifier }
