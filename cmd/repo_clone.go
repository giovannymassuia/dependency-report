package cmd

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/giovannymassuia/dependency-report/cmd/utils"
	"github.com/giovannymassuia/dependency-report/internal/repo"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository",
	Long:  `Clone a repository from specified source.`,
	Run: func(cmd *cobra.Command, args []string) {

		if true {
			utils.PrintError("This command is disabled for now!")
			return
		}

		provider, err := utils.GetRepositoryProvider(cmd)
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		name, err := flags.GetRequiredFlag(cmd, "name")
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		err = provider.CloneRepository(name)
		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		fmt.Printf("Repository [%s] cloned successfully in the `%s` folder\n", name, repo.TempDir)
	},
}

func init() {
	repoCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().String("name", "", "Repository name")
}
