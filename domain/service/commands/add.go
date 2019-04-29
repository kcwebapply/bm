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
	tags, _ := util.GetTerminalInput("tags (input few tags by ',')")

	tagSize := len(strings.Split(tags, ","))
	if tagSize > 3 {
		fmt.Println("tag size error. you can't put more than 3 tags on your bookmark.")
		os.Exit(0)
	}

	newPage := add(url, tags)

	view.PrintAdd(newPage)
}

func add(url string, tags string) model.Page {
	title, content, err := http.GetContent(url)
	if err != nil {
		fmt.Println("error get http contents.")
		os.Exit(0)
	}
	newPage := model.Page{URL: url, Title: *title, Tags: tags, Content: *content}
	repository.AddPage(newPage)
	return newPage
}
