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

// isGitRepo returns true if the given path is a git repository.
func isGitRepo(path string) bool {
	fullPath := filepath.Join(path, ".git")
	_, err := os.Stat(fullPath)
	return !os.IsNotExist(err)
}

// isExists returns true if the given path exists.
func isExists(path string) bool {
	_, err := os.Stat(path)
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
