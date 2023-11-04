package repo

import "fmt"

type RepoProvider interface {
	ListRepositories() ([]Repository, error)
	CloneRepository(name string) error
	ListRepositoryDependencies(name string) error
}

type DefaultRepoProvider struct {
}

func (DefaultRepoProvider) ListRepositoryDependencies(name string) ([]Repository, error) {
	return nil, fmt.Errorf("not implemented")
}

func (DefaultRepoProvider) CloneRepository(name string) error {
	return fmt.Errorf("not implemented")
}

func (DefaultRepoProvider) ListRepositories() error {
	return fmt.Errorf("not implemented")
}

type Repository struct {
	Name    string `json:"name"`
	Private bool   `json:"is_private"`
}
