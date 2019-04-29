package commands

import (
	"github.com/codegangsta/cli"
)

// Commands return cli-command setting list .
func Commands() []cli.Command {

	return []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "register web page on bm.",
			Action:  Add,
		},
		{
			Name:    "open",
			Aliases: []string{"p"},
			Usage:   "open url ",
			Action:  OpenPage,
		},
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "view bookmark list.",
			Action:  Ls,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t,tag",
					Value: "",
					Usage: "search tag",
				},

				cli.StringFlag{
					Name:  "s,search",
					Value: "",
					Usage: "content search parameter (receive word argument)",
				},
			},
		},
		{
			Name:    "tags",
			Aliases: []string{"t"},
			Usage:   "tagList",
			Action:  GetTags,
		},
		{
			Name:    "rm",
			Aliases: []string{"r"},
			Usage:   "delete bookmark ",
			Action:  Rm,
		},
	}
}