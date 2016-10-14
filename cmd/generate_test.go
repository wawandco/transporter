package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestGenerate(t *testing.T) {
	setupTestingEnv()
	base := os.Getenv("TRANS_TESTING_FOLDER")
	os.RemoveAll(base)
	os.Mkdir(base, generatedFilePermissions)

	var context = cli.Context{}
	Init(&context)
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
