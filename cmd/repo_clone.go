package cmd

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/giovannymassuia/dependency-report/cmd/utils"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository",
	Long:  `Clone a repository from specified source.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Printf("Repository [%s] cloned successfully in the `./temp` folder\n", name)
	},
}

func init() {
	repoCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().String("name", "", "Repository name")
}
