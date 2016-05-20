package utils

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/yaml.v1"
)

type TestConfig map[string]map[string]string

type TestMigration struct {
	Identifier  int64
	UpCommand   string
	DownCommand string
}

var testingTables = []string{
	"other_table",
	"down_table",
	"transporter_migrations",
}

func ClearTestTables(driver string) {
	conn, _ := BuildTestingConnection(driver)
	defer conn.Close()

	for _, t := range testingTables {
		conn.Exec("DROP TABLE IF EXISTS  " + t)
	}
}

func ClearTestMigrations() {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.RemoveAll(filepath.Join(base, "db"))
}

func SetupTestingFolders(driver string) {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.Mkdir(filepath.Join(base), 0777)
	os.Mkdir(filepath.Join(base, "db"), 0777)
	os.Mkdir(filepath.Join(base, "db", "migrations"), 0777)

	BuildTestConfigFile(base, driver)
}

func BuildTestConfigFile(base, driver string) {
	url := os.Getenv(strings.ToUpper(driver) + "_DATABASE_URL")

	data := TestConfig{
		"development": {
			"url":    url,
			"driver": driver,
		},
	}

	d, _ := yaml.Marshal(&data)

	filepath := filepath.Join(base, "db", "config.yml")
	ioutil.WriteFile(filepath, d, 0777)
}

func GenerateMigrationFiles(migs []TestMigration) {
	for _, mig := range migs {
		GenerateMigrationFile(mig)
	}
}

func GenerateMigrationFile(mig TestMigration) {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	buff := bytes.NewBufferString("")
	tmpl, _ := template.New("test_migration").Parse(MigrationTemplate)
	_ = tmpl.Execute(buff, mig)

	fileName := strconv.FormatInt(mig.Identifier, 10) + "_a.go"
	path := filepath.Join(base, "db", "migrations", fileName)
	ioutil.WriteFile(path, buff.Bytes(), 0777)
}

func BuildTestingConnection(driver string) (*sql.DB, error) {
	url := os.Getenv(strings.ToUpper(driver) + "_DATABASE_URL")
	return sql.Open(driver, url)
}
