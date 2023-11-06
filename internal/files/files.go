package files

import (
	"os"
	"path/filepath"
)

// deleteFolder deletes the given folder.
func deleteFolder(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

// isFolderAGitRepo returns true if the given path is a git repository.
func isFolderAGitRepo(path string) bool {
	gitDir := filepath.Join(path, ".git")
	_, err := os.Stat(gitDir)
	return !os.IsNotExist(err)
}

type File struct {
	FileName   string
	FilePath   string
	FolderName string
	FolderPath string
}

// findFiles returns a list of files that match the given file name searching recursively from the given root path.
// Returns empty list if no files are found.
func findFiles(rootPath, fileName string) ([]string, error) {
	return nil, nil
}

// isFilePresentInFolder returns true if the given file name is present in the given folder path.
func isFilePresentInFolder(folderPath, fileName string) bool {
	return false
}
