package cmd

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
	"github.com/wawandco/transporter/transporter"
)

var sampleMigrations = []transporter.Migration{
	transporter.Migration{
		Identifier: transporter.MigrationIdentifier(),
		Up: func(tx *sql.Tx) {
			tx.Exec("Create table tests_table (a varchar);")
		},
		Down: func(tx *sql.Tx) {
			tx.Exec("Drop table tests_table;")
		},
	},

	transporter.Migration{
		Identifier: transporter.MigrationIdentifier(),
		Up: func(tx *sql.Tx) {
			tx.Exec("ALTER table tests_table ADD COLUMN other varchar(20);")
		},
		Down: func(tx *sql.Tx) {
			tx.Exec("ALTER table tests_table DROP COLUMN other;")
		},
	},
}

func exists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}

func buildConnectionFromConfig() (*sql.DB, error) {

	if os.Getenv("TRANS_TESTING_FOLDER") != "" {
		return sql.Open("postgres", "user=transporter dbname=transporter sslmode=disable")
	}

	//TODO: Real connection from config
	return nil, nil
}

func cleanTables() {
	db, _ := buildConnectionFromConfig()
	db.Exec("DROP TABLE IF EXISTS  " + transporter.MigrationsTable + ";")
	db.Exec("DROP TABLE IF EXISTS other_table ;")
	db.Exec("DROP TABLE IF EXISTS down_table ;")
}

func writeTemplateToFile(path string, t *template.Template, data interface{}) (string, error) {
	f, e := os.Create(path)
	if e != nil {
		return "", e
	}
	defer f.Close()

	e = t.Execute(f, data)
	if e != nil {
		return "", e
	}

	return f.Name(), nil
}

func replaceInFile(file, base, replacement string) {

	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if line == base {
			lines[i] = replacement
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}
}