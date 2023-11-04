package cmd

import (
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Repository commands",
	Long: `Manage repositories from various sources. For example:
	
	List all repositories from Bitbucket:
	$ dependency-report repoCmds list --bitbucket
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(repoCmd)

	repoCmd.PersistentFlags().Bool(flags.ProviderBitbucket, false, "Use Bitbucket as source")
	repoCmd.PersistentFlags().Bool(flags.ProviderGithub, false, "Use Github as source")

	// bitbucket specific flags
	repoCmd.PersistentFlags().String(flags.BitbucketWorkspace, "", "Bitbucket workspace [required]")
	repoCmd.PersistentFlags().String(flags.BitbucketAppPassword, "", "Bitbucket app password")
	repoCmd.PersistentFlags().String(flags.BitbucketOAuthKey, "", "Bitbucket OAuth client ID [requires --bb-oauth-secret]")
	repoCmd.PersistentFlags().String(flags.BitbucketOAuthSecret, "", "Bitbucket OAuth client secret [requires --bb-oauth-key]")
}
