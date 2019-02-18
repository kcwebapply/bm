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

	searchFlag := []cli.Flag{
		cli.StringFlag{
			Name:  "t,tag",
			Value: "",
			Usage: "search tag",
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
			Flags:   searchFlag,
		},
		{
			Name:    "tags",
			Aliases: []string{"t"},
			Usage:   "tagList",
			Action:  commands.GetTags,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
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
