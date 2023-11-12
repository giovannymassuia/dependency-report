package internal

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
)

type Node struct{}

func (n Node) Scan() ([]dependencies.ReportModel, error) {
	return nil, fmt.Errorf("scan method not implemented for node")
}
