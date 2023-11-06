package managers

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
)

type Maven struct{}

func (m Maven) Scan() ([]dependencies.ReportModel, error) {
	return nil, fmt.Errorf("scan method not implemented for maven")
}
