package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/crufter/copyrecur"
)

func WriteTemplateToFile(path string, t *template.Template, data interface{}) (string, error) {
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

func ReplaceInFile(file, base, replacement string) {

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

func CopyMigrationFilesTo(tempFolder string) []string {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	copyrecur.CopyFile(filepath.Join(base, "db", "config.yml"), filepath.Join(tempFolder, "config.yml"))
	commandArgs := []string{"run"}

	files, _ := ioutil.ReadDir(filepath.Join(base, "db", "migrations"))
	for _, file := range files {
		commandArgs = append(commandArgs, filepath.Join(tempFolder, file.Name()))
		copyrecur.CopyFile(filepath.Join(base, "db", "migrations", file.Name()), filepath.Join(tempFolder, file.Name()))
		ReplaceInFile(filepath.Join(tempFolder, file.Name()), "package migrations", "package main")
	}
	return commandArgs
}
