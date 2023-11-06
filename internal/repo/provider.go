package repo

import (
	"fmt"
)

const TempDir = ".dependency-reports-temp"

type Repository struct {
	Name    string `json:"name"`
	Private bool   `json:"is_private"`
}

type RepoProvider interface {
	ListRepositories() ([]Repository, error)
	CloneRepository(name string) error
	ListRepositoryDependencies(all bool, name string) ([]Project, error)
}

type DefaultRepoProvider struct {
}

func (DefaultRepoProvider) ListRepositoryDependencies(bool, string) error {
	return fmt.Errorf("not implemented")
}

func (DefaultRepoProvider) CloneRepository(string) error {
	return fmt.Errorf("not implemented")
}

func (DefaultRepoProvider) ListRepositories() ([]Repository, error) {
	return nil, fmt.Errorf("not implemented")
}
