package managers

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
	return "DROP TABLE IF EXISTS " + tableName + ";"
}

//CreateMigrationsTableQuery is the implementation of how to create migraitons table for this particular manager.
func (man *PostgreSQLManager) CreateMigrationsTableQuery(tableName string) string {
	return "CREATE TABLE IF NOT EXISTS " + tableName + " ( identifier decimal NOT NULL );"
}

//LastMigrationQuery is the implementation of how to return the last runt migration for this particular manager.
func (man *PostgreSQLManager) LastMigrationQuery(tableName string) string {
	return "SELECT max(identifier) FROM " + tableName + ";"
}
