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
	case 4:
		result, err = gitemail.GetCommitsInRepo(os.Args[1], os.Args[2], os.Args[3])
	case 3:
		result, err = gitemail.GetCommitsInUser(os.Args[1], os.Args[2])
	default:
		log.Fatal("Wrong number of arguments")
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result {
		fmt.Printf("%s\n", item)
	}
}
