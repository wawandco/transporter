package managers

//PostgreSQLManager is the manager for postgresql DBMS
type MySQLManager struct{}

func (man *MySQLManager) AllMigrationsQuery(tableName string) string {
	return "SELECT * FROM " + tableName
}

func (man *MySQLManager) DeleteMigrationQuery(tableName string, identifier string) string {
	return "DELETE FROM " + tableName + " WHERE identifier = " + identifier
}

func (man *MySQLManager) AddMigrationQuery(tableName string, identifier string) string {
	return "INSERT INTO " + tableName + " ( identifier ) VALUES (" + identifier + ")"
}

func (man *MySQLManager) DropMigrationsTableQuery(tableName string) string {
	return "DROP TABLE IF EXISTS " + tableName
}

func (man *MySQLManager) CreateMigrationsTableQuery(tableName string) string {
	return "CREATE TABLE IF NOT EXISTS " + tableName + " ( identifier BIGINT )"
}

func (man *MySQLManager) LastMigrationQuery(tableName string) string {
	return "SELECT MAX(identifier) FROM " + tableName
}
