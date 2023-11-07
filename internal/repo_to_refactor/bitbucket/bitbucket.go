package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/giovannymassuia/dependency-report/internal/repo_to_refactor"
	"io"
	"net/http"
	"os/exec"
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
	Values  []repo_to_refactor.Repository `json:"values"`
	Page    int                           `json:"page"`
	PageLen int                           `json:"pagelen"`
	Size    int                           `json:"size"`
	Next    string                        `json:"next"`
	Prev    string                        `json:"previous"`
}

func (p *Provider) ListRepositories() ([]repo_to_refactor.Repository, error) {

	response, err := getRepositoriesApi(p, 1)
	if err != nil {
		return nil, err
	}

	var repositories []repo_to_refactor.Repository
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
	var auth string
	if p.auth.oauth != nil {
		accessToken, err := p.getAccessToken()
		if err != nil {
			return err
		}
		auth = fmt.Sprintf("x-token-auth:%s", accessToken)
	} else if p.auth.appPassword != "" {
		auth = fmt.Sprintf("%s:%s", p.workspace, p.auth.appPassword)
	} else {
		return fmt.Errorf("missing authentication")
	}

	// check if temp with repository name already exists
	if repo_to_refactor.GitRepositoryExists(fmt.Sprintf("%s/%s", repo_to_refactor.TempDir, name)) {
		return fmt.Errorf("repository already cloned")
	}

	repoURL := fmt.Sprintf("https://%s@bitbucket.org/%s/%s.git", auth, p.workspace, name)

	cmd := exec.Command("git", "clone", repoURL, fmt.Sprintf("%s/%s", repo_to_refactor.TempDir, name))
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (p *Provider) ListRepositoryDependencies(all bool, name string) ([]repo_to_refactor.Project, error) {
	if all {
		return nil, fmt.Errorf("repo_to_refactor dependencies with --all: not implemented")
	} else {
		if name == "" {
			return nil, fmt.Errorf("missing repository name")
		}
		return repo_to_refactor.ScanRepository(name, p.CloneRepository)
	}
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
