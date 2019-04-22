package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/page"
	view "github.com/kcwebapply/bm/view"
)

// DeletePage delete pagedata
func DeletePage(c *cli.Context) {

	id := c.Args().Get(0)
	if id == "" {
		os.Exit(0)
	}
	page := deletePage(id)

	view.PrintSavePage(page)
}

func deletePage(id string) page.Page {
	allPages := readLines()
	writer := getFileCleanWriter(fileName)
	defer writer.Flush()

	var deletePage page.Page

	for _, page := range allPages {
		if strconv.Itoa(page.ID) == id {
			deletePage = page
			continue
		}
		writer.Write(([]byte)(page.String()))
	}

	if err := os.Remove(contentPath + "/" + id + ".txt"); err != nil {
		fmt.Println(err)
	}

	return deletePage
}
