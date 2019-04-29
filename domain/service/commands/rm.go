package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// Rm delete pagedata
func Rm(c *cli.Context) {
	id := c.Args().Get(0)
	if id == "" {
		os.Exit(0)
	}
	page := rm(id)

	view.PrintRm(page)
}

func rm(id string) model.Page {
	err := repository.RemovePage(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// temp
	return model.Page{}
}
