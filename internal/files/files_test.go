package files

import (
	test_utils "github.com/giovannymassuia/dependency-report/test/testutils"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestFindFiles(t *testing.T) {
	modulePath := test_utils.GetModulePath()
	fullPath := filepath.Join(modulePath, "test/data")
	files, err := FindFiles(fullPath, "test_pom.xml")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(files))
	assert.Equal(t, "test_pom.xml", files[0].FileName)
	assert.Equal(t, filepath.Join(fullPath, "/files_to_find/sub_folder_to_find/test_pom.xml"), files[0].FilePath)
	assert.Equal(t, "sub_folder_to_find", files[0].FolderName)
	assert.Equal(t, filepath.Join(fullPath, "/files_to_find/sub_folder_to_find"), files[0].FolderPath)
	assert.Equal(t, "test_pom.xml", files[1].FileName)
	assert.Equal(t, filepath.Join(fullPath, "/files_to_find/test_pom.xml"), files[1].FilePath)
	assert.Equal(t, "files_to_find", files[1].FolderName)
	assert.Equal(t, filepath.Join(fullPath, "/files_to_find"), files[1].FolderPath)
}

func TestCheckIfPathExists(t *testing.T) {
	modulePath := test_utils.GetModulePath()
	fullPath := filepath.Join(modulePath, "test/data/temp/fake_repo")
	testGitFile, err := os.Create(path.Join(fullPath, ".git"))
	if err != nil {
		t.Errorf("Error creating .git file: %v", err)
	}

	assert.True(t, IsGitRepo(fullPath))

	// remove .git file
	err = os.Remove(testGitFile.Name())
	assert.Nil(t, err)
	assert.False(t, IsGitRepo(fullPath))
}

func TestIsExists(t *testing.T) {
	modulePath := test_utils.GetModulePath()
	fullPath := filepath.Join(modulePath, "test/data/maven_repo")
	assert.True(t, IsExists(fullPath, "pom.xml"))
}

func TestDeleteFolder(t *testing.T) {
	modulePath := test_utils.GetModulePath()
	fullPath := filepath.Join(modulePath, "test/data/temp/to_beRemoved")
	err := os.MkdirAll(fullPath, os.ModePerm)
	assert.Nil(t, err)
	assert.True(t, isFolderExist(fullPath))

	err = DeleteFolder(fullPath)
	assert.Nil(t, err)
	assert.False(t, isFolderExist(fullPath))
}

func TestDeleteFolder_WhenFolderDoesNotExist(t *testing.T) {
	modulePath := test_utils.GetModulePath()
	fullPath := filepath.Join(modulePath, "test/data/temp/to_beRemoved")
	assert.False(t, isFolderExist(fullPath))

	err := DeleteFolder(fullPath)
	assert.Nil(t, err)
	assert.False(t, isFolderExist(fullPath))
}

func isFolderExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
