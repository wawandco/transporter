package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/crufter/copyrecur"
)

//WriteTemplateToFile writes a template result to a file.
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

//ReplaceInFile Replaces a text inside a file, this is usefull to replace the migrations package for the
//main when copying the sources to run Up and Down.
func ReplaceInFile(file, base, replacement string) {

	input, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	reg, err := regexp.Compile(base)

	for i, line := range lines {
		if reg.MatchString(line) {
			lines[i] = replacement
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)

	if err != nil {
		log.Fatalln(err)
	}
}

//CopyMigrationFilesTo copies the migration files into a specific folder, this is useful to run the Up and Down commands.
func CopyMigrationFilesTo(tempFolder string) []string {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	copyrecur.CopyFile(filepath.Join(base, "db", "config.yml"), filepath.Join(tempFolder, "config.yml"))
	commandArgs := []string{"run"}

	files, _ := ioutil.ReadDir(filepath.Join(base, "db", "migrations"))
	for _, file := range files {
		commandArgs = append(commandArgs, filepath.Join(tempFolder, file.Name()))
		copyrecur.CopyFile(filepath.Join(base, "db", "migrations", file.Name()), filepath.Join(tempFolder, file.Name()))
		ReplaceInFile(filepath.Join(tempFolder, file.Name()), "^package migrations$", "package main")
		ReplaceInFile(filepath.Join(tempFolder, file.Name()), "^(.*)github.com/wawandco/transporter/core(.*)$", "  transporter \"github.com/wawandco/transporter/core\"")
	}
	return commandArgs
}
