package repo

import (
	"os"
	"path/filepath"
)

func GitRepositoryExists(path string) bool {
	gitDir := filepath.Join(path, ".git")
	_, err := os.Stat(gitDir)
	return !os.IsNotExist(err)
}
