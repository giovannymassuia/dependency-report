package dependencies

import "fmt"

type ManagerInterface interface {
	// Scan scans a repository and returns a list of dependencies.ReportModel
	Scan(path string) ([]ReportModel, error)
}

type ManagerDefault struct {
}

func (m ManagerDefault) Scan() ([]ReportModel, error) {
	return nil, fmt.Errorf("scan method not implemented")
}
