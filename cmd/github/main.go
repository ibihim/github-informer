package main

import (
	"fmt"

	"github.com/ibihim/github-informer/pkg/config"
	"github.com/ibihim/github-informer/pkg/github"

	gh "github.com/google/go-github/v42/github"
)

func main() {
	c := config.Flags()
	if c.Verbose {
		fmt.Println(c)
	}

	client := gh.NewClient(nil)

	if c.Query == "" {
		q, err := github.ToQuery(c)
		if err != nil {
			fmt.Println(err)
			return
		}

		c.Query = q
	}

	i, err := github.Query(client, c.Query)
	if err != nil {
		fmt.Println(err)
		return
	}

	github.PrintIssues(i)
}
