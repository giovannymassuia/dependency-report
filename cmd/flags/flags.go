package flags

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	ProviderBitbucket = "bitbucket"
	ProviderGithub    = "github"

	BitbucketAppPassword = "bb-app-password"
	BitbucketOAuthKey    = "bb-oauth-key"
	BitbucketOAuthSecret = "bb-oauth-secret"
	BitbucketWorkspace   = "bb-workspace"

	DependencyManager = "manager"
)

func GetRequiredFlag(cmd *cobra.Command, flag string) (string, error) {
	value, err := cmd.Flags().GetString(flag)
	if err != nil {
		return "", err
	}

	if value == "" {
		return "", fmt.Errorf("%s is required", FormatFlagOutput(flag))
	}

	return value, nil
}

func FormatFlagOutput(flag string) string {
	return fmt.Sprintf("\033[37m--%s\033[0m", flag)
}
