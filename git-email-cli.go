package main

import (
	"fmt"
	"github.com/jpcano/git-email/lib"
	"log"
	"os"
)

func main() {
	var err error
	var result []string

	switch len(os.Args) {
	case 3:
		result, err = gitemail.GetCommitsInUser(os.Args[1], os.Args[2])
	case 4:
		result, err = gitemail.GetCommitsInRepo(os.Args[1], os.Args[2], os.Args[3])
	default:
		fmt.Printf("usage: git-email <GitHub User> <email>\n")
		fmt.Printf("       git-email <GitHub User> <Repository> <email>\n")
		os.Exit(-1)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result {
		fmt.Printf("%s\n", item)
	}
}
