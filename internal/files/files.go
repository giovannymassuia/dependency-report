package files

import (
	"os"
	"path/filepath"
)

// DeleteFolder deletes the given folder.
func DeleteFolder(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

// IsGitRepo returns true if the given path is a git repository.
func IsGitRepo(path string) bool {
	fullPath := filepath.Join(path, ".git")
	_, err := os.Stat(fullPath)
	return !os.IsNotExist(err)
}

// IsExists returns true if the given path exists.
func IsExists(path, file string) bool {
	_, err := os.Stat(filepath.Join(path, file))
	return !os.IsNotExist(err)
}

type File struct {
	FileName   string
	FilePath   string
	FolderName string
	FolderPath string
}

// FindFiles returns a list of files that match the given file name searching recursively from the given root path.
// Returns empty list if no files are found.
func FindFiles(rootPath, fileName string) ([]File, error) {
	var files []File
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == fileName {
			files = append(files, File{
				FileName:   info.Name(),
				FilePath:   path,
				FolderName: filepath.Base(filepath.Dir(path)),
				FolderPath: filepath.Dir(path),
			})
		}
		return nil
	})
	return files, err
}
