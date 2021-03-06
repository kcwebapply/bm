package commands

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/util"
	"github.com/kcwebapply/bm/view"
)

// Rm delete pagedata
func Rm(c *cli.Context) {
	id := c.Args().Get(0)

	err := rm(id)

	if err != nil {
		util.LoggingError(err.Error())
	}

	view.PrintRm(id)
}

func rm(id string) error {
	err := repository.RemovePage(id)
	return err
}
