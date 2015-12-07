package managers

//DatabaseManager is an interface defined for the database drivers we will use,
//it basically contains the operations we will need to do on each DBMS, and some
//times need specific sql code.
type DatabaseManager interface {
	AllMigrationsQuery(tableName string) string
	DeleteMigrationQuery(tableName string, identifier string) string
	AddMigrationQuery(tableName string, identifier string) string
	DropMigrationsTableQuery(tableName string) string
	CreateMigrationsTableQuery(tableName string) string
	LastMigrationQuery(tableName string) string
	DropTableQuery(tableName string) string
	CreateTableQuery(tableName string, tableColumns Table) string
	AddColumnQuery(tableName string, columnName string, columnType string) string
	DropColumnQuery(tableName string, columnName string) string
	ChangeColumnTypeQuery(tableName string, columnName string, newType string) string
	RenameColumnQuery(tableName string, columnName string, newName string) string
	RenameTableQuery(tableName string, newName string) string
}

// Table is the type to define tables, basically is the column and its type.
type Table map[string]string
