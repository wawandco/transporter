package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	setupTestingEnv()

	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.Mkdir(filepath.Join(base), generatedFilePermissions)
	os.Mkdir(filepath.Join(base, "db"), generatedFilePermissions)
	os.Mkdir(filepath.Join(base, "db", "migrations"), generatedFilePermissions)

	var context = cli.Context{}
	Generate(&context)

	files, _ := ioutil.ReadDir(filepath.Join(base, "db", "migrations"))
	assert.Equal(t, 1, len(files))

	for _, file := range files {
		filepath := filepath.Join(base, "db", "migrations", file.Name())
		content, _ := ioutil.ReadFile(filepath)
		assert.Contains(t, string(content), "Identifier:")
	}
}

func TestGenerateWithoutFolder(t *testing.T) {
	setupTestingEnv()
	var context = cli.Context{}

	assert.NotPanics(t, func() {
		Generate(&context)
	})
}
