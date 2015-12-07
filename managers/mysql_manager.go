package managers

import "strings"

//MySQLManager is the manager for mysql DBMS
type MySQLManager struct{}

//AllMigrationsQuery is the implementation of how to get all migrations for this particular manager.
func (man *MySQLManager) AllMigrationsQuery(tableName string) string {
	return "SELECT * FROM " + tableName
}

//DeleteMigrationQuery is the implementation of how to delete a migration for this particular manager.
func (man *MySQLManager) DeleteMigrationQuery(tableName string, identifier string) string {
	return "DELETE FROM " + tableName + " WHERE identifier = " + identifier
}

//AddMigrationQuery is the implementation of how to add a migration for this particular manager.
func (man *MySQLManager) AddMigrationQuery(tableName string, identifier string) string {
	return "INSERT INTO " + tableName + " ( `identifier` ) VALUES (" + identifier + ")"
}

//DropMigrationsTableQuery is the implementation of how to drop migraitons table for this particular manager.
func (man *MySQLManager) DropMigrationsTableQuery(tableName string) string {
	return man.DropTableQuery(tableName)
}

//CreateMigrationsTableQuery is the implementation of how to create migraitons table for this particular manager.
func (man *MySQLManager) CreateMigrationsTableQuery(tableName string) string {
	return "CREATE TABLE IF NOT EXISTS " + tableName + " ( `identifier` BIGINT )"
}

//LastMigrationQuery is the implementation of how to return the last runt migration for this particular manager.
func (man *MySQLManager) LastMigrationQuery(tableName string) string {
	return "SELECT MAX(identifier) FROM " + tableName
}

//DropTableQuery is the implementation of how to drop table for this particular manager.
func (man *MySQLManager) DropTableQuery(tableName string) string {
	return "DROP TABLE IF EXISTS " + tableName
}

func (man *MySQLManager) CreateTableQuery(tableName string, tableColumns Table) string {
	query := "CREATE TABLE " + tableName + " ("
	columns := []string{}
	for column, tipe := range tableColumns {
		columns = append(columns, column+" "+tipe+"")
	}

	columnsString := strings.Join(columns, ", ")
	return query + columnsString + ")"
}

//AddColumnQuery is the implementation of how to add a column table for this particular manager.
func (man *MySQLManager) AddColumnQuery(tableName string, columnName string, columnType string) string {
	return "ALTER TABLE " + tableName + " ADD COLUMN " + columnName + " " + columnType
}

//DropColumnQuery is the implementation of how to drop a column for this particular manager.
func (man *MySQLManager) DropColumnQuery(tableName string, columnName string) string {
	return "ALTER TABLE " + tableName + " DROP COLUMN " + columnName
}

//ChangeColumnTypeQuery is the implementation of how to change column type for this particular manager.
func (man *MySQLManager) ChangeColumnTypeQuery(tableName string, columnName string, newType string) string {
	return "ALTER TABLE " + tableName + " CHANGE COLUMN " + columnName + " " + newType
}

//RenameColumnQuery is the implementation of how to change column name for this particular manager.
func (man *MySQLManager) RenameColumnQuery(tableName string, columnName string, newName string) string {
	return "SIGNAL SQLSTATE '78000' SET MESSAGE_TEXT = 'MySQL seems not to support table column renaming without passing the column type.';"
}

//RenameTableQuery is the implementation of how to change table name for this particular manager.
func (man *MySQLManager) RenameTableQuery(tableName string, newName string) string {
	return "RENAME TABLE " + tableName + " TO " + newName
}
