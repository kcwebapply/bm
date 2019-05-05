package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/infrastructure/http"
	"github.com/kcwebapply/bm/view"
)

// Add saves pagedata
func Add(c *cli.Context) {

	url := c.Args().Get(0)

	var tagList = []string{}
	for i := 1; i <= 3; i++ {
		var tag = c.Args().Get(i)
		if tag != "" {
			tagList = append(tagList, tag)
		}
	}

	newPage, err := add(url, strings.Join(tagList, ","))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	view.PrintAdd(*newPage)
}

func add(url string, tags string) (*model.Page, error) {
	title, content, err := http.GetContent(url)
	if err != nil {
		return nil, err
	}
	newPage := model.Page{URL: url, Title: *title, Tags: tags, Content: *content}
	err = repository.AddPage(newPage)
	if err != nil {
		return nil, err
	}
	return &newPage, err
}
