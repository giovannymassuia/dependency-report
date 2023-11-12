package managers

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
)

type ManagerInterface interface {
	// Scan scans a repository and returns a list of dependencies.ReportModel
	Scan(path string) ([]dependencies.ReportModel, error)
}

type ManagerDefault struct {
}

func (m ManagerDefault) Scan() ([]dependencies.ReportModel, error) {
	return nil, fmt.Errorf("scan method not implemented")
}
