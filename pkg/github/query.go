package github

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ibihim/github-informer/pkg/config"

	"github.com/google/go-github/v42/github"
)

const isOpen = "is:open"

var errNoFlags = errors.New("set query or a couple of other flags instead")

func Query(c *github.Client, q string) ([]*github.Issue, error) {
	isr, _, err := c.Search.Issues(context.Background(), q, nil)
	if err != nil {
		return nil, err
	}

	return isr.Issues, nil
}

func ToQuery(c config.Conf) (string, error) {
	var builder strings.Builder

	builder.WriteString(isOpen)

	// the follow up arg must contain a whitespace as prefix
	if c.User != "" {
		if _, err := builder.WriteString(" author:"); err != nil {
			return "", err
		}
		if _, err := builder.WriteString(c.User); err != nil {
			return "", err
		}
	}

	// the follow up arg must contain a whitespace as prefix
	if c.Kind != "" {
		if _, err := builder.WriteString(" is:"); err != nil {
			return "", err
		}
		if _, err := builder.WriteString(c.Kind); err != nil {
			return "", err
		}
	}

	// the follow up arg must contain a whitespace as prefix
	if c.Repo != "" {
		if _, err := builder.WriteString("repo:"); err != nil {
			return "", err
		}
		if _, err := builder.WriteString(c.Repo); err != nil {
			return "", err
		}
	}

	if result := builder.String(); result != isOpen {
		return result, nil
	}

	return "", errNoFlags
}

func PrintIssues(is []*github.Issue) {
	var currentUser string

	for _, i := range is {
		if currentUser != *i.User.Login {
			currentUser = *i.User.Login

			fmt.Printf(`
user: %s
issues:`,
				currentUser)
		}

		PrintIssue(i)
	}
}

func PrintIssue(i *github.Issue) {
	fmt.Printf(`
- title: %s
  updated: %s
  comments: %d
  url: %s`,
		*i.Title,
		i.UpdatedAt.Format("2006-01-02"),
		*i.Comments,
		*i.HTMLURL,
	)
}
