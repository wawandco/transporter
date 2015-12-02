package managers

//DatabaseManager is an interface defined for the database drivers we will use.
type DatabaseManager interface {
	AllMigrationsQuery(tableName string) string
	DeleteMigrationQuery(tableName string, identifier string) string
	AddMigrationQuery(tableName string, identifier string) string
	DropMigrationsTableQuery(tableName string) string
	CreateMigrationsTableQuery(tableName string) string
	LastMigrationQuery(tableName string) string
}