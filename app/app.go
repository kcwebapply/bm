package app

import (
	"log"

	"github.com/codegangsta/cli"
	commands "github.com/kcwebapply/bm/domain/service/commands"
)

var appName = "bm"
var version = "0.0.1"

// InitApp method is for initializing and  getting App settings.
func InitApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "github.com/kcwebapply/bm"
	app.Version = version
	// command routing.
	app.Commands = commands.Commands()

	app.Before = func(c *cli.Context) error {
		err := ArgumentSizeValidator(c)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	return app
}
