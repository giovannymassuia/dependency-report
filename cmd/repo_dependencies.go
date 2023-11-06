package cmd

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/giovannymassuia/dependency-report/cmd/utils"

	"github.com/spf13/cobra"
)

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

		all, _ := cmd.Flags().GetBool("all")

		var name string
		if !all {
			name, err = flags.GetRequiredFlag(cmd, "name")
			if err != nil {
				utils.PrintError(err.Error())
				return
			}
		}

		repositories, err := provider.ListRepositoryDependencies(all, name)
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		output, _ := cmd.Flags().GetString("output")
		err = utils.PrintRepositories(repositories, output)
		if err != nil {
			utils.PrintError(err.Error())
			return
		}
	},
}

func init() {
	repoCmd.AddCommand(dependenciesCmd)

	dependenciesCmd.Flags().StringP("name", "n", "", "Repository name")
	dependenciesCmd.Flags().BoolP("all", "a", false, "Scan all repositories")
	dependenciesCmd.Flags().String("output", "print", "Output format print|json|csv")
}
