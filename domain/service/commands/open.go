package commands

import (
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/repository"
	open_cmd "github.com/skratchdot/open-golang/open"
)

// Open open bookmark website
func Open(c *cli.Context) {
	id := c.Args().Get(0)
	open(id)
}

func open(id string) {
	pages, _ := repository.GetPages()
	for _, page := range pages {
		if id == strconv.Itoa(page.ID) {
			open_cmd.Run(page.URL)
		}
	}
}
