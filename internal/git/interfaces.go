package git

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
)

type ProviderInterface interface {
	Clone() error
	List() ([]string, error)
	Dependencies(repoName string) ([]dependencies.ReportModel, error)
}

type ProviderDefault struct {
}

func (ProviderDefault) Clone() error {
	return fmt.Errorf("clone method not implemented")
}

func (ProviderDefault) List() ([]string, error) {
	return nil, fmt.Errorf("list method not implemented")
}

func (ProviderDefault) Dependencies(repoName string) ([]dependencies.ReportModel, error) {
	return nil, fmt.Errorf("dependencies method not implemented")
}
