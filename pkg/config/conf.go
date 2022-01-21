package config

import (
	"flag"
	"fmt"
)

type Conf struct {
	User    string
	Kind    string
	Repo    string
	Query   string
	Verbose bool
}

func (c Conf) String() string {
	return fmt.Sprintf(`
Config:
	- user: %s
	- kind: %s
	- repo: %s
	- query: %s
	- verbose: %t
	`,
		c.User,
		c.Kind,
		c.Repo,
		c.Query,
		c.Verbose,
	)
}

func Flags() Conf {
	var c Conf

	flag.StringVar(&c.User, "u", "ibihim", "user name to check")
	flag.StringVar(&c.Kind, "t", "pr", "type to search for")
	flag.StringVar(&c.Repo, "r", "", "filter to given repository")
	flag.StringVar(&c.Query, "q", "", "query by hand, if set ignores all other flags")
	flag.BoolVar(&c.Verbose, "v", false, "prints out additional data")

	flag.Parse()

	return c
}
