package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/wawandco/transporter/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func init() {
	gopath := os.Getenv("GOPATH")
	testingDir := filepath.Join(gopath, "src", "github.com", "wawandco", "transporter", "testing")
	err := os.RemoveAll(testingDir)

	if err != nil {
		fmt.Println(err)
	}

	os.Setenv("TRANS_TESTING_FOLDER", testingDir)
}

func TestInit(t *testing.T) {
	setupTestingEnv()
	context := cli.Context{}
	Init(&context)
	base := os.Getenv("TRANS_TESTING_FOLDER")

	files := []string{
		filepath.Join(base, "db"),
		filepath.Join(base, "db", "migrations"),
		filepath.Join(base, "db", "config.yml"),
	}

	for _, file := range files {
		isThere, _ := exists(file)
		assert.True(t, isThere)
	}

	// //Content
	content, _ := ioutil.ReadFile(files[2])
	assert.Contains(t, string(content), "database")
	assert.Contains(t, string(content), "url")
	assert.Contains(t, string(content), "driver")

}

func TestInitExistingFolder(t *testing.T) {
	setupTestingEnv()
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.Mkdir(filepath.Join(base, "db"), generatedFilePermissions)

	context := cli.Context{}
	Init(&context)
	isThere, _ := exists(filepath.Join(base, "db", "migrations"))
	assert.False(t, isThere)

}

func setupTestingEnv() {
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.RemoveAll(filepath.Join(base, "db"))
}
