package gitemail

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

// Commits

var commits_url string = "https://api.github.com/repos/%s/%s/commits"

type Commits []*CommitData
	
type CommitData struct {
	Commit  *Commit
	HTMLURL string `json:"html_url"`
}

type Commit struct {
	Author *Author
}

type Author struct {
	Email string
}

// Repos

var repos_url string = "https://api.github.com/users/%s/repos"

type Repos []*RepoData

type RepoData struct {
	Name string 
}

func Fetch(query string) (io.ReadCloser, error) {
	res, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("GitHub API request failed: %s", res.Status)
	}
	return res.Body, nil
}

func FetchCommits(user, repo string) (Commits, error) {
	url := fmt.Sprintf(commits_url, user, repo)
	rc, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	var result Commits
	if err := json.NewDecoder(rc).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func FetchRepos(user string) (Repos, error) {
	url := fmt.Sprintf(repos_url, user)
	rc, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	var result Repos
	if err := json.NewDecoder(rc).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetCommitsInRepo(user, repo, email string) ([]string, error) {
	var result []string
	commits, err := FetchCommits(user, repo)
	if err != nil {
		return nil, err
	}
	for _, commit := range commits {
		if commit.Commit.Author.Email == email {
			result = append(result, commit.HTMLURL)
		}
	}
	return result, nil
}

func GetCommitsInUser(user, email string) ([]string, error) {
	var result []string
	var commits []string
	repos, err := FetchRepos(user)
	if err != nil {
		return nil, err
	}
	for _, repo := range repos {
		commits, err = GetCommitsInRepo(user, repo.Name, email)
		if err != nil {
			return nil, err
		}
		for _, commit := range commits {
			result = append(result, commit)
		}
	}
	return result, nil
}
