package app

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/commands"
)

var appName = "bm"
var version = "0.0.1"

// InitApp method is for initializing and  getting App settings.
func InitApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "github.com/kcwebapply/bm"
	app.Version = version

	lsFlag := []cli.Flag{
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
	}

	// command routing.
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "register web page on bm.",
			Action:  commands.SavePage,
		},
		{
			Name:    "open",
			Aliases: []string{"p"},
			Usage:   "open url ",
			Action:  commands.OpenPage,
		},
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "view bookmark list.",
			Action:  commands.GetAllPages,
			Flags:   lsFlag,
		},
		{
			Name:    "tags",
			Aliases: []string{"t"},
			Usage:   "tagList",
			Action:  commands.GetTags,
		},
		{
			Name:    "rm",
			Aliases: []string{"r"},
			Usage:   "delete bookmark ",
			Action:  commands.DeletePage,
		},
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	return app
}
