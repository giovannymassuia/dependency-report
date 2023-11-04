package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/dependencies"
	"github.com/giovannymassuia/dependency-report/internal/repo"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	git_http "github.com/go-git/go-git/v5/plumbing/transport/http"
	"io"
	"net/http"
	"os"
)

const (
	bitbucketAPIBaseURL    = "https://api.bitbucket.org/2.0"
	bitbucketOAuthEndpoint = "https://bitbucket.org/site/oauth2/access_token"
)

func New(workspace, appPassword, oauthKey, oauthSecret string) (*Provider, error) {
	var oauth *OAuth
	// both oauthKey and oauthSecret must be set if one of them is set
	if oauthKey != "" || oauthSecret != "" {
		if oauthKey == "" || oauthSecret == "" {
			return nil, fmt.Errorf("missing OAuth key or secret")
		}
		oauth = &OAuth{
			key:    oauthKey,
			secret: oauthSecret,
		}
	}

	if workspace == "" {
		return nil, fmt.Errorf("missing workspace")
	}

	p := Provider{
		workspace: workspace,
		auth: Auth{
			appPassword: appPassword,
			oauth:       oauth,
		},
	}

	return &p, nil
}

type repositoriesResponse struct {
	Values  []repo.Repository `json:"values"`
	Page    int               `json:"page"`
	PageLen int               `json:"pagelen"`
	Size    int               `json:"size"`
	Next    string            `json:"next"`
	Prev    string            `json:"previous"`
}

func (p *Provider) ListRepositories() ([]repo.Repository, error) {

	response, err := getRepositoriesApi(p, 1)
	if err != nil {
		return nil, err
	}

	var repositories []repo.Repository
	repositories = append(repositories, response.Values...)

	// while next is not empty, get next page and append to values
	for response.Next != "" {
		response, err = getRepositoriesApi(p, response.Page+1)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, response.Values...)
	}

	return repositories, nil
}

func getRepositoriesApi(p *Provider, page int) (*repositoriesResponse, error) {
	// get repositories
	url := fmt.Sprintf("%s/repositories/%s", bitbucketAPIBaseURL, p.workspace)

	if page > 1 {
		url = fmt.Sprintf("%s?page=%d", url, page)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if err := p.setAuth(req); err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get repositories, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response repositoriesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (p *Provider) CloneRepository(name string) error {
	// Configure the authentication
	var auth transport.AuthMethod
	// if oauth is set, use it as bearer token
	if p.auth.oauth != nil {
		accessToken, err := p.getAccessToken()
		if err != nil {
			return err
		}
		auth = &git_http.TokenAuth{
			Token: accessToken,
		}
	} else if p.auth.appPassword != "" {
		auth = &git_http.BasicAuth{
			Username: p.workspace,
			Password: p.auth.appPassword,
		}
	}

	repoURL := fmt.Sprintf("https://bitbucket.org/%s/%s.git", p.workspace, name)

	// Cloning the repository
	_, err := git.PlainClone(fmt.Sprintf("temp/%s", name), false, &git.CloneOptions{
		Auth:     auth,
		URL:      repoURL,
		Progress: os.Stdout, // Display progress
	})
	if err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	}

	return nil
}

func (p *Provider) ListRepositoryDependencies(name string) error {
	pomFiles, err := dependencies.FindPomFiles(fmt.Sprintf("temp/%s", name))
	if err != nil {
		return err
	}

	//dependencies.ReadDependencyTree("temp/leferrante-api/deps.txt")
	for _, pomFile := range pomFiles {

		project, err := dependencies.ReadPomFile(pomFile)
		if err != nil {
			return err
		}

		fmt.Printf("Project: %s\n", project.ArtifactId)
		fmt.Printf("Version: %s\n", project.Version)
		fmt.Printf("Parent: %s:%s:%s\n", project.Parent.GroupId, project.Parent.ArtifactId, project.Parent.Version)
		for _, dependency := range project.Dependencies.Dependency {
			fmt.Printf("  - %s:%s:%s:%s\n", dependency.GroupId, dependency.ArtifactId, dependency.Version, dependency.Scope)
		}

	}

	return nil
}

func (p *Provider) setAuth(req *http.Request) error {
	if p.auth.oauth != nil {
		// if oauth is set, use it as bearer token
		accessToken, err := p.getAccessToken()
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	} else if p.auth.appPassword != "" {
		// use basic auth
		req.SetBasicAuth(p.workspace, p.auth.appPassword)
	}
	return nil
}

func (p *Provider) getAccessToken() (string, error) {

	// validate auth
	if p.auth.appPassword == "" && p.auth.oauth == nil {
		return "", fmt.Errorf("missing authentication")
	}

	client := &http.Client{}

	data := "grant_type=client_credentials"
	req, err := http.NewRequest("POST", bitbucketOAuthEndpoint, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(p.auth.oauth.key, p.auth.oauth.secret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse OAuthTokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil
}
