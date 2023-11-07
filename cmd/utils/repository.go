package utils

import (
	"fmt"
	"github.com/giovannymassuia/dependency-report/cmd/flags"
	"github.com/giovannymassuia/dependency-report/internal/repo_to_refactor"
	"github.com/giovannymassuia/dependency-report/internal/repo_to_refactor/bitbucket"
	"github.com/spf13/cobra"
)

func GetRepositoryProvider(cmd *cobra.Command) (repo_to_refactor.RepoProvider, error) {
	bitbucketFlag, _ := cmd.Flags().GetBool(flags.ProviderBitbucket)
	githubFlag, _ := cmd.Flags().GetBool(flags.ProviderGithub)

	if bitbucketFlag {
		workspace, err := flags.GetRequiredFlag(cmd, flags.BitbucketWorkspace)
		if err != nil {
			return nil, err
		}
		oauthKey, _ := cmd.Flags().GetString(flags.BitbucketOAuthKey)
		oauthSecret, _ := cmd.Flags().GetString(flags.BitbucketOAuthSecret)
		appPassword, _ := cmd.Flags().GetString(flags.BitbucketAppPassword)

		if oauthKey == "" && oauthSecret != "" {
			return nil, fmt.Errorf("%s is required if %s is set", flags.FormatFlagOutput(flags.BitbucketOAuthKey), flags.FormatFlagOutput(flags.BitbucketOAuthSecret))
		} else if oauthKey != "" && oauthSecret == "" {
			return nil, fmt.Errorf("%s is required if %s is set", flags.FormatFlagOutput(flags.BitbucketOAuthSecret), flags.FormatFlagOutput(flags.BitbucketOAuthKey))
		}

		return bitbucket.New(workspace, appPassword, oauthKey, oauthSecret)
	} else if githubFlag {
		//provider = github.New()
		return nil, fmt.Errorf("Github is not yet supported")
	} else {
		return nil, fmt.Errorf("no provider specified. Use %s or %s", flags.FormatFlagOutput(flags.ProviderBitbucket), flags.FormatFlagOutput(flags.ProviderGithub))
	}
}
