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

func fetch_commits (user, repo, email string) (Commits ,error){
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

func main() {
	result, err := fetch_commits(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result {
		fmt.Printf("%s\n", item.HTMLURL)
	}
}
