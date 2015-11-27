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
	Database map[string]string
}

// Init creates needed files/folders for transporter to work correctly.
func Init(ctx *cli.Context) {
	base := ""
	if os.Getenv("TRANS_TESTING_FOLDER") != "" {
		base = os.Getenv("TRANS_TESTING_FOLDER")
		os.Mkdir(base, generatedFilePermissions)
	}

	if isThere, _ := exists(filepath.Join(base, "db")); isThere {
		log.Println("| db folder already exists")
		return
	}

	os.Mkdir(filepath.Join(base, "db"), generatedFilePermissions)
	os.Mkdir(filepath.Join(base, "db", "migrations"), generatedFilePermissions)

	data := map[string]string{
		"url":    "$DATABASE_URL",
		"driver": "$DATABASE_DRIVER",
	}

	config := Config{data}
	d, _ := yaml.Marshal(&config)

	filepath := filepath.Join(base, "db", "config.yml")
	ioutil.WriteFile(filepath, d, generatedFilePermissions)
}
