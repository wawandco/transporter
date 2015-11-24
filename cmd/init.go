package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v1"
)

const (
	generatedFilePermissions = 0777
)

type Config struct {
	Database map[string]string
}

func Init() {
	base := ""
	if os.Getenv("TRANS_TESTING_FOLDER") != "" {
		base = os.Getenv("TRANS_TESTING_FOLDER")
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
	fmt.Println(filepath)
	err := ioutil.WriteFile(filepath, d, generatedFilePermissions)
	fmt.Println(err)
}
