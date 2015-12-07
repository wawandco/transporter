package managers

import "strings"

//PostgreSQLManager is the manager for postgresql DBMS
type PostgreSQLManager struct{}

//AllMigrationsQuery is the implementation of how to get all migrations for this particular manager.
func (man *PostgreSQLManager) AllMigrationsQuery(tableName string) string {
	return "SELECT * FROM " + tableName + ";"
}

//DeleteMigrationQuery is the implementation of how to delete a migration for this particular manager.
func (man *PostgreSQLManager) DeleteMigrationQuery(tableName string, identifier string) string {
	return "DELETE FROM " + tableName + " WHERE identifier = " + identifier + ";"
}

//AddMigrationQuery is the implementation of how to add a migration for this particular manager.
func (man *PostgreSQLManager) AddMigrationQuery(tableName string, identifier string) string {
	return "INSERT INTO " + tableName + " ( identifier ) VALUES (" + identifier + ");"
}

//DropMigrationsTableQuery is the implementation of how to drop migraitons table for this particular manager.
func (man *PostgreSQLManager) DropMigrationsTableQuery(tableName string) string {
	return man.DropTableQuery(tableName)
}

//CreateMigrationsTableQuery is the implementation of how to create migraitons table for this particular manager.
func (man *PostgreSQLManager) CreateMigrationsTableQuery(tableName string) string {
	return "CREATE TABLE IF NOT EXISTS " + tableName + " ( identifier decimal NOT NULL );"
}

//LastMigrationQuery is the implementation of how to return the last runt migration for this particular manager.
func (man *PostgreSQLManager) LastMigrationQuery(tableName string) string {
	return "SELECT max(identifier) FROM " + tableName + ";"
}

//DropTableQuery is the implementation of how to drop table for this particular manager.
func (man *PostgreSQLManager) DropTableQuery(tableName string) string {
	return "DROP TABLE " + tableName + ";"
}

// CreateTableQuery returns the query to create a table based on the table name and the table structure.
func (man *PostgreSQLManager) CreateTableQuery(tableName string, tableColumns Table) string {
	query := "CREATE TABLE " + tableName + " ("
	columns := []string{}
	for column, tipe := range tableColumns {
		columns = append(columns, column+" "+tipe+"")
	}

	columnsString := strings.Join(columns, ", ")
	return query + columnsString + ")"
}

// AddColumnQuery returns the query to add a column into a table.
func (man *PostgreSQLManager) AddColumnQuery(tableName string, columnName string, columnType string) string {
	return "ALTER TABLE " + tableName + " ADD COLUMN " + columnName + " " + columnType + ";"
}

//DropColumnQuery is the implementation of how to drop a column for this particular manager.
func (man *PostgreSQLManager) DropColumnQuery(tableName string, columnName string) string {
	return "ALTER TABLE " + tableName + " DROP COLUMN " + columnName
}

//ChangeColumnTypeQuery is the implementation of how to change column type for this particular manager.
func (man *PostgreSQLManager) ChangeColumnTypeQuery(tableName string, columnName string, newType string) string {
	return "ALTER TABLE " + tableName + " ALTER COLUMN " + columnName + " SET DATA TYPE " + newType
}

//RenameColumnQuery is the implementation of how to change column name for this particular manager.
func (man *PostgreSQLManager) RenameColumnQuery(tableName string, columnName string, newName string) string {
	return "ALTER TABLE " + tableName + " RENAME COLUMN " + columnName + " TO " + newName + ";"
}

//RenameTableQuery is the implementation of how to change table name for this particular manager.
func (man *PostgreSQLManager) RenameTableQuery(tableName string, newName string) string {
	return "ALTER TABLE " + tableName + " RENAME TO " + newName
}
