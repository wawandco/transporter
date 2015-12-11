package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/wawandco/transporter/managers"
	"github.com/wawandco/transporter/transporter"
)

//MainData is the data passed to generate the main.go when running Up and Down.
type MainData struct {
	TempDir     string
	Environment string
}

var sampleMigrations = []transporter.Migration{
	transporter.Migration{
		Identifier: transporter.MigrationIdentifier(),
		Up: func(tx *transporter.Tx) {
			tx.Exec("Create table tests_table (a varchar);")
		},
		Down: func(tx *transporter.Tx) {
			tx.Exec("Drop table tests_table;")
		},
	},

	transporter.Migration{
		Identifier: transporter.MigrationIdentifier(),
		Up: func(tx *transporter.Tx) {
			tx.Exec("ALTER table tests_table ADD COLUMN other varchar(20);")
		},
		Down: func(tx *transporter.Tx) {
			tx.Exec("ALTER table tests_table DROP COLUMN other;")
		},
	},
}

var mans = map[string]managers.DatabaseManager{
	"postgres": &managers.PostgreSQLManager{},
	"mysql":    &managers.MySQLManager{},
}

func exists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}

func setupTestingEnv() {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.Mkdir(base, 0777)
}

func runTempFiles(commandArgs []string) {
	cmd := exec.Command("go", commandArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if e := cmd.Run(); e != nil {
		log.Fatal("`go run` failed: ", e)
	}
}

func buildTempFolder() string {
	temp, e := ioutil.TempDir("", "transporter")
	if e != nil {
		log.Fatal(e)
	}
	return temp
}
