package core

import (
	"database/sql"

	"github.com/wawandco/transporter/managers"
)

//Tx is our implementation of the sql.Tx, we did this to make the sql errors
//visible to transporter user when the databse operation fails.
type Tx struct {
	*sql.Tx
	Manager managers.DatabaseManager
	err     error
}

//Exec is a wrapper to the Exec function inside sql.Tx.
func (tx *Tx) Exec(query string) (sql.Result, error) {
	return tx.execAndSaveErr(query)
}

// DropTable function drops a particular table from the database based on the driver implementation.
func (tx *Tx) DropTable(name string) (sql.Result, error) {
	query := tx.Manager.DropTableQuery(name)
	return tx.execAndSaveErr(query)
}

// CreateTable allows us to create tables by passing a table definition and the table name.
func (tx *Tx) CreateTable(name string, columns managers.Table) (sql.Result, error) {
	query := tx.Manager.CreateTableQuery(name, columns)
	return tx.execAndSaveErr(query)
}

// AddColumn allows us to add columns to tables by passing column name and type.
func (tx *Tx) AddColumn(name string, columnName string, tipe string) (sql.Result, error) {
	query := tx.Manager.AddColumnQuery(name, columnName, tipe)
	return tx.execAndSaveErr(query)
}

// DropColumn allows us to drop columns by passing column name on table
func (tx *Tx) DropColumn(name string, columnName string) (sql.Result, error) {
	query := tx.Manager.DropColumnQuery(name, columnName)
	return tx.execAndSaveErr(query)
}

// ChangeColumnType provides an interface to change columns types.
func (tx *Tx) ChangeColumnType(name string, columnName string, newType string) (sql.Result, error) {
	query := tx.Manager.ChangeColumnTypeQuery(name, columnName, newType)
	return tx.execAndSaveErr(query)
}

// RenameColumn provides an interface to change columns names.
func (tx *Tx) RenameColumn(name string, columnName string, newName string) (sql.Result, error) {
	query := tx.Manager.RenameColumnQuery(name, columnName, newName)
	return tx.execAndSaveErr(query)
}

// RenameTable provides an interface to change columns names.
func (tx *Tx) RenameTable(name string, newName string) (sql.Result, error) {
	query := tx.Manager.RenameTableQuery(name, newName)
	return tx.execAndSaveErr(query)
}

func (tx *Tx) execAndSaveErr(query string) (sql.Result, error) {
	result, err := tx.Tx.Exec(query)

	if err != nil {
		tx.err = err
	}

	return result, err
}
