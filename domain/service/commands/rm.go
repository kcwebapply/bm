package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// Rm delete pagedata
func Rm(c *cli.Context) {
	id := c.Args().Get(0)
	if id == "" {
		os.Exit(0)
	}

	err := rm(id)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	view.PrintRm(id)
}

func rm(id string) error {
	err := repository.RemovePage(id)
	return err
}
