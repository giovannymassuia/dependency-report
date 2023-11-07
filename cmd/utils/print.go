package utils

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/repo_to_refactor"
)

func PrintError(msg string) {
	// in red
	fmt.Println("\033[31mError:\033[0m", msg)
}

func PrintRepositories(repositories []repo_to_refactor.Project, output string) error {
	if output == "json" {
		//printRepositoriesJson(repositories)
		return fmt.Errorf("json not implemented")
	} else if output == "csv" {
		//printRepositoriesCsv(repositories)
		return fmt.Errorf("csv not implemented")
	} else {
		return printRepositories(repositories)
	}
}

func printRepositories(projects []repo_to_refactor.Project) error {

	for _, project := range projects {
		fmt.Printf("Project: %s\n", project.ArtifactId)
		fmt.Printf("Version: %s\n", project.Version)
		fmt.Printf("Parent: %s:%s:%s\n", project.Parent.GroupId, project.Parent.ArtifactId, project.Parent.Version)

		fmt.Printf("Dependencies:\n")
		for _, dependency := range project.Dependencies.Dependency {
			fmt.Printf("  - %s:%s:%s:%s\n", dependency.GroupId, dependency.ArtifactId, dependency.Version, dependency.Scope)
		}

		fmt.Printf("Dependency Management:\n")
		for _, dependency := range project.DependencyManagement.Dependencies.Dependency {
			fmt.Printf("  - %s:%s:%s:%s\n", dependency.GroupId, dependency.ArtifactId, dependency.Version, dependency.Scope)
		}
	}

	return nil
}
