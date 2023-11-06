package files

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestIsFolderAGitRepo(t *testing.T) {
	modulePath := getModulePath()
	fullPath := filepath.Join(modulePath, "test/data/temp/fake_repo")
	testGitFile, err := os.Create(filepath.Join(fullPath, ".git"))
	if err != nil {
		t.Errorf("Error creating .git file: %v", err)
	}

	assert.True(t, isFolderAGitRepo(fullPath))

	// remove .git file
	err = os.Remove(testGitFile.Name())
	assert.Nil(t, err)
	assert.False(t, isFolderAGitRepo(fullPath))
}

func TestDeleteFolder(t *testing.T) {
	modulePath := getModulePath()
	fullPath := filepath.Join(modulePath, "test/data/temp/to_beRemoved")
	err := os.MkdirAll(fullPath, os.ModePerm)
	assert.Nil(t, err)
	assert.True(t, isFolderExist(fullPath))

	err = deleteFolder(fullPath)
	assert.Nil(t, err)
	assert.False(t, isFolderExist(fullPath))
}

func TestDeleteFolder_WhenFolderDoesNotExist(t *testing.T) {
	modulePath := getModulePath()
	fullPath := filepath.Join(modulePath, "test/data/temp/to_beRemoved")
	assert.False(t, isFolderExist(fullPath))

	err := deleteFolder(fullPath)
	assert.Nil(t, err)
	assert.False(t, isFolderExist(fullPath))
}

func isFolderExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getModulePath() string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
		return ""
	}
	path := out.String()
	// remove \n
	path = path[:len(path)-1]
	return path
}
