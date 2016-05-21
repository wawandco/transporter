package core

import (
	"database/sql"
	"errors"
	"log"
	"math"
	"sort"
	"strconv"
	"sync"

	//Driver require to be loaded this way.
	"github.com/go-sql-driver/mysql"
	//Driver require to be loaded this way.
	_ "github.com/lib/pq"

	"github.com/wawandco/transporter/managers"
	"gopkg.in/yaml.v1"
)

var mu sync.Mutex

func init() {
	mu.Lock()
	defer mu.Unlock()

	if !driverRegistered("mysql") {
		sql.Register("mysql", &mysql.MySQLDriver{})
	}

	if !driverRegistered("mariadb") {
		sql.Register("mariadb", &mysql.MySQLDriver{})
	}
}

var migrations []Migration
var manager managers.DatabaseManager
var databaseManagers = map[string]managers.DatabaseManager{
	"postgres": &managers.PostgreSQLManager{},
	"mysql":    &managers.MySQLManager{},
	"mariadb":  &managers.MariaDBManager{},
}

const (
	// MigrationsTable is the Db table where we will store migrations
	MigrationsTable = "transporter_migrations"
)

// Add function adds a migration to the migrations array
// So Transporter verifies if the migration is already completed or runs
// the desired migration as needed.
//
// Register checks if migration attempted to be registered id is unique.
func Add(m Migration) {
	migrations = append(migrations, m)
}

// MigrationsTableExists returns true if the table for the migrations already exists.
func MigrationsTableExists(db *sql.DB) bool {
	query := manager.AllMigrationsQuery(MigrationsTable)
	_, err := db.Query(query)
	return err == nil
}

//RunAllMigrationsUp runs pending migrations and stores migration on the migrations table.
func RunAllMigrationsUp(db *sql.DB) {

	//1. Migrations table exists? -> Create if needed
	if !MigrationsTableExists(db) {
		CreateMigrationsTable(db)
	}

	sort.Sort(ByIdentifier(migrations))

	version := DatabaseVersion(db)
	if len(migrations) > 0 && version == migrations[len(migrations)-1].GetID() {
		log.Println("| No migrations to run, DB is on latest version. (" + version + ")")
		return
	}

	var err error
	for _, migration := range migrations {
		// Check if migration is on the database already
		if migration.Pending(db) {
			err = RunMigrationUp(db, &migration)

			if err != nil {
				log.Println("| Could not complete your migration (" + migration.GetID() + "), please check your SQL.")
				log.Println("| " + err.Error())
				break
			}
		}
	}

	version = DatabaseVersion(db)
	if err == nil && version != "" {
		log.Println("| Done, new database version is " + version)
	}
}

// RunMigrationUp runs a single migration up and if success it saves the
// Migration identifier on the migrations table.
func RunMigrationUp(db *sql.DB, m *Migration) error {
	tx, err := dbTransaction(db)
	if err != nil {
		return errors.New("Could not open a db transaction.")
	}

	if m.Up != nil {
		log.Printf("| Running Migration %s UP", m.GetID())

		m.Up(tx)
		migerr := tx.err
		err = tx.Commit()

		if err != nil || migerr != nil {
			return migerr
		}

		query := manager.AddMigrationQuery(MigrationsTable, m.GetID())
		_, err = db.Exec(query)
		return nil
	}

	return errors.New("Migration doesnt have Up function defined.")
}

// RunMigrationDown Runs one migration down and if successful it removes the migration from
// the migrations table.
func RunMigrationDown(db *sql.DB, m *Migration) error {
	tx, err := dbTransaction(db)
	if err != nil {
		return errors.New("Could not begin a transaction.")
	}

	if m.Down != nil {
		log.Printf("| Running Migration %s DOWN", m.GetID())
		m.Down(tx)
		migerr := tx.err
		err = tx.Commit()

		if err != nil || migerr != nil {
			return migerr
		}
		query := manager.DeleteMigrationQuery(MigrationsTable, m.GetID())
		_, err = db.Exec(query)
		return nil
	}

	return errors.New("Migration (" + m.GetID() + ") doesn't have Down function defined.")
}

//RunOneMigrationDown Run down last migration that have completed.
func RunOneMigrationDown(db *sql.DB) {
	if !MigrationsTableExists(db) {
		CreateMigrationsTable(db)
	}

	sort.Sort(ByIdentifier(migrations))
	identifier := DatabaseVersion(db)
	id, err := strconv.Atoi(identifier)

	if err != nil {
		log.Println("Sorry, there is no migration to run back")
	}

	for _, mig := range migrations {
		if mig.Identifier-int64(id) == 0 {
			log.Println("| Running " + mig.GetID() + " back")
			err := RunMigrationDown(db, &mig)
			if err == nil {
				version := DatabaseVersion(db)
				if version != "" {
					log.Println("| Done, new database version is " + version)
				} else {
					log.Println("| Done, All existing migrations down.")
				}
			} else {
				log.Println("| Could not rollback your migration (" + mig.GetID() + "), please check your SQL.")
				log.Println("| " + err.Error())
			}
			break
		}
	}

}

//DatabaseVersion returns the latest database version.
func DatabaseVersion(db *sql.DB) string {
	query := manager.LastMigrationQuery(MigrationsTable)
	rows, err := db.Query(query)
	if err != nil || rows == nil {
		log.Println(err)
		return strconv.FormatInt(math.MaxInt64, 10)
	}

	var identifier string
	for rows.Next() {
		rows.Scan(&identifier)
		break
	}

	return identifier
}

//SetManager allows external entities like testing to set the driver as needed.
func SetManager(man managers.DatabaseManager) {
	manager = man
}

//DBConnection Returns a DB connection from the yml config file
func DBConnection(ymlFile []byte, environment string) (*sql.DB, error) {
	var connData map[string]map[string]string
	err := yaml.Unmarshal([]byte(ymlFile), &connData)

	if err != nil {
		return nil, err
	}

	if connData[environment] == nil {
		err = errors.New("Environment [" + environment + "] does not exist in your config.yml")
		return nil, err
	}

	manager = databaseManagers[connData[environment]["driver"]]
	return sql.Open(connData[environment]["driver"], connData[environment]["url"])
}

func dbTransaction(db *sql.DB) (*Tx, error) {
	sqlTx, err := db.Begin()
	if err != nil {
		log.Println("| Error, could not initialize transaction")
	}

	tx := &Tx{
		Tx:      sqlTx,
		Manager: manager,
		err:     nil,
	}

	return tx, err
}
