package cmd

// func init() {
// 	utils.ClearTestTables()
// 	setupTestingEnv()
// }

//
// func TestInitCommand(t *testing.T) {
// 	setupTestingEnv()
// 	utils.ClearTestMigrations()
//
// 	context := cli.Context{}
// 	Init(&context)
// 	base := os.Getenv("TRANS_TESTING_FOLDER")
//
// 	files := []string{
// 		filepath.Join(base, "db"),
// 		filepath.Join(base, "db", "migrations"),
// 		filepath.Join(base, "db", "config.yml"),
// 	}
//
// 	for _, file := range files {
// 		isThere, _ := exists(file)
// 		assert.True(t, isThere)
// 	}
//
// 	// //Content
// 	content, _ := ioutil.ReadFile(files[2])
// 	assert.Contains(t, string(content), "development")
// 	assert.Contains(t, string(content), "url")
// 	assert.Contains(t, string(content), "driver")
//
// }
//
// func TestInitExistingFolder(t *testing.T) {
// 	setupTestingEnv()
// 	base := os.Getenv("TRANS_TESTING_FOLDER")
//
// 	os.RemoveAll(base)
// 	os.Mkdir(filepath.Join(base), generatedFilePermissions)
// 	os.Mkdir(filepath.Join(base, "db"), generatedFilePermissions)
//
// 	context := cli.Context{}
// 	Init(&context)
// 	isThere, _ := exists(filepath.Join(base, "db", "migrations"))
// 	assert.False(t, isThere)
// }
