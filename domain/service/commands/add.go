package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/infrastructure/http"
	"github.com/kcwebapply/bm/util"
	"github.com/kcwebapply/bm/view"
)

// Add saves pagedata
func Add(c *cli.Context) {

	url := c.Args().Get(0)

	title, _ := util.GetTerminalInput("title")

	tags, _ := util.GetTerminalInput("tags (input few tags by ',')")

	tagSize := len(strings.Split(tags, ","))
	if tagSize > 3 {
		fmt.Println("tag size error. you can't put more than 3 tags on your bookmark.")
		os.Exit(0)
	}

	content, err := http.GetContent(url)
	if err != nil {
		fmt.Println("error get http contents.")
		os.Exit(0)
	}

	newPage := add(url, title, tags, *content)

	view.PrintAdd(newPage)
}

func add(url string, title string, tags string, content string) model.Page {
	allPages, _ := repository.GetPages()
	pageSize := len(allPages)
	var newID = 1
	if pageSize > 0 {
		newID = allPages[pageSize-1].ID + 1
	}

	newPage := model.Page{ID: newID, URL: url, Title: title, Tags: tags, Content: content}
	repository.AddPage(newPage)
	return newPage
}
