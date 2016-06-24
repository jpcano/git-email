package main 
	
import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"os"
)

// parse json
type jsonUser struct {
	Name string `json:"login"`
	Blog string `json:"blog"`
}

var commits_url string = "https://api.github.com/repos/%s/%s/commits"
// var repos_url string = "https://api.github.com/users/%s/repos"

type Commits []*CommitData

type CommitData struct {
	Commit *Commit
	HTMLURL string `json:"html_url"`
}

type Commit struct {
	Author *Author
}

type Author struct {
	Email string
}

func FetchCommits (user, repo string) (Commits, error){
	// request http api
	url := fmt.Sprintf(commits_url, user, repo)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("GitHub API request failed: %s", res.Status)
	}

	var result Commits
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()
	return result, nil
}

func GetCommitsByEmail (user, repo, email string) ([]string, error) {
	var result []string
	commits, err := FetchCommits(user, repo)
	if err != nil {
		return nil, err
	}
	for _, commit := range commits {
		// fmt.Printf("%s\n", commit.HTMLURL)
		if commit.Commit.Author.Email == email {
			result = append(result, commit.HTMLURL)
		}
	}
	return result, nil
}

func main() {
	result, err := GetCommitsByEmail(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result {
		fmt.Printf("%s\n", item)
	}
}
