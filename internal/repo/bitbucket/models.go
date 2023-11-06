package bitbucket

import "github.com/giovannymassuia/dependency-report/internal/repo"

type Provider struct {
	repo.DefaultRepoProvider
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
