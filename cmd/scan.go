package cmd

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/giovannymassuia/dependency-report/cmd/utils"
	"github.com/giovannymassuia/dependency-report/internal/dependencies/managers"
	"github.com/spf13/cobra"
	"os"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan Dependencies",
	Long:  `Scan dependencies for a local project. Execute it in the root directory of the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if flag not maven, return not supported
		managerFlag := cmd.Flag(flags.DependencyManager)
		if managerFlag.Value.String() != "maven" {
			utils.PrintError(("This dependency manager is not supported yet!"))
			return
		}

		// if flag maven, execute maven scan
		// use gpwd
		currentPath, err := os.Getwd()
		if err != nil {
			utils.PrintError(err.Error())
			return
		}
		result, err := managers.NewMaven().Scan(currentPath)
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		for _, report := range result {
			fmt.Println(report.Project.Name)
			// deps
			for _, dep := range report.Dependencies {
				fmt.Println(dep.Name + " " + dep.Version)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.PersistentFlags().StringP(flags.DependencyManager, "m", "maven", "Dependency manager (maven|gradle|npm)")
	scanCmd.MarkPersistentFlagRequired(flags.DependencyManager)
}
