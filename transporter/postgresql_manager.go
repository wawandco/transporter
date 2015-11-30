package transporter

//PostgreSQLManager is the manager for postgresql DBMS
type PostgreSQLManager struct{}

func (man *PostgreSQLManager) AllMigrationsQuery(tableName string) string {
	return "SELECT * FROM " + tableName + ";"
}

func (man *PostgreSQLManager) DeleteMigrationQuery(tableName string, identifier string) string {
	return "DELETE FROM " + tableName + " WHERE identifier = " + identifier + ";"
}

func (man *PostgreSQLManager) AddMigrationQuery(tableName string, identifier string) string {
	return "INSERT INTO " + tableName + " ( identifier ) VALUES ('" + identifier + "') ;"
}

func (man *PostgreSQLManager) DropMigrationsTableQuery(tableName string) string {
	return "DROP TABLE IF EXISTS " + tableName + ";"
}

func (man *PostgreSQLManager) CreateMigrationsTableQuery(tableName string) string {
	return "CREATE TABLE IF NOT EXISTS " + tableName + ";"
}

func (man *PostgreSQLManager) LastMigrationQuery(tableName string) string {
	return "SELECT max(identifier) FROM " + tableName + ";"
}
