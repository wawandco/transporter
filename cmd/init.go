package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/gopkg.in/yaml.v1"
)

const (
	generatedFilePermissions = 0777
)

//Config holds configuration values for transporter
type Config struct {
	environments map[string]map[string]string
}

// Init creates needed files/folders for transporter to work correctly.
func Init(ctx *cli.Context) {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.Mkdir(base, generatedFilePermissions)

	if isThere, _ := exists(filepath.Join(base, "db")); isThere {
		log.Println("| db folder already exists")
		return
	}

	os.Mkdir(filepath.Join(base, "db"), generatedFilePermissions)
	os.Mkdir(filepath.Join(base, "db", "migrations"), generatedFilePermissions)

	data := map[string]map[string]string{
		"development": {
			"url":    "$DATABASE_URL",
			"driver": "$DATABASE_DRIVER",
		},
	}

	d, _ := yaml.Marshal(data)
	filepath := filepath.Join(base, "db", "config.yml")
	ioutil.WriteFile(filepath, d, generatedFilePermissions)
}
