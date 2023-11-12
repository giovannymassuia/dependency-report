package dependencies

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies/internal/managers"
)

func ManagerFactory(manager string) (ManagerInterface, error) {
	switch manager {
	case "maven":
		return managers.NewMaven(), nil
	default:
		return nil, fmt.Errorf("manager %s not implemented", manager)
	}
}
