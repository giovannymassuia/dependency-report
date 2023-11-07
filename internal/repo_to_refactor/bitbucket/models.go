package bitbucket

import "github.com/giovannymassuia/dependency-report/internal/repo_to_refactor"

type Provider struct {
	repo_to_refactor.DefaultRepoProvider
	workspace string
	auth      Auth
}

type Auth struct {
	appPassword string
	oauth       *OAuth
}

type OAuth struct {
	key    string
	secret string
}

type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}
