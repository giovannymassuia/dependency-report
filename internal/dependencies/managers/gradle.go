package managers

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
)

type Gradle struct{}

func (g Gradle) Scan() ([]dependencies.ReportModel, error) {
	return nil, fmt.Errorf("scan method not implemented for gradle")
}
