package cmd

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/giovannymassuia/dependency-report/cmd/utils"

	"github.com/spf13/cobra"
)

// dependenciesCmd represents the dependencies command
var dependenciesCmd = &cobra.Command{
	Use:   "dependencies",
	Short: "List repository dependencies",
	Long:  `List repository dependencies from specified source.`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, err := utils.GetRepositoryProvider(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}

		name, err := flags.GetRequiredFlag(cmd, "name")
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		err = provider.ListRepositoryDependencies(name)
		if err != nil {
			fmt.Printf("Error reading dependencies from repository: %v\n", err)
			return
		}
	},
}

func init() {
	repoCmd.AddCommand(dependenciesCmd)

	dependenciesCmd.Flags().String("name", "", "Repository name")
}
