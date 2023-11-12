package managers

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies/managers/internal"
)

func ManagerFactory(manager string) (ManagerInterface, error) {
	switch manager {
	case "maven":
		return internal.NewMaven(), nil
	default:
		return nil, fmt.Errorf("manager %s not implemented", manager)
	}
}
