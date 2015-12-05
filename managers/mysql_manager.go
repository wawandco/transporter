package managers

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
	return "DROP TABLE IF EXISTS " + tableName
}

//CreateMigrationsTableQuery is the implementation of how to create migraitons table for this particular manager.
func (man *MySQLManager) CreateMigrationsTableQuery(tableName string) string {
	return "CREATE TABLE IF NOT EXISTS " + tableName + " ( `identifier` BIGINT )"
}

//LastMigrationQuery is the implementation of how to return the last runt migration for this particular manager.
func (man *MySQLManager) LastMigrationQuery(tableName string) string {
	return "SELECT MAX(identifier) FROM " + tableName
}
