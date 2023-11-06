package cmd

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all repositories",
	Long:  `List all repositories from specified source.`,
	Run: func(cmd *cobra.Command, args []string) {
		provider, err := utils.GetRepositoryProvider(cmd)
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		repos, err := provider.ListRepositories()
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		if len(repos) == 0 {
			fmt.Println("No repositories found")
			return
		}

		for _, repo := range repos {
			public := ""
			if !repo.Private {
				public = " (public)"
			}
			fmt.Printf("%s%s\n", repo.Name, public)
		}

	},
}

func init() {
	repoCmd.AddCommand(listCmd)
}
