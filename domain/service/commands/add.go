package commands

import (
	"fmt"
	"net/url"
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

	urlString := c.Args().Get(0)

	if _, err := url.Parse(urlString); err != nil {
		util.LoggingError(fmt.Sprintf("failed to parse URL %q: %s", urlString, err))
	}

	var tagList = c.Args().Tail()

	newPage, err := add(urlString, tagList)

	if err != nil {
		util.LoggingError(fmt.Sprintf("add commands error : %s", err))
	}

	view.PrintAdd(*newPage)
}

func add(url string, tags []string) (*model.Page, error) {
	title, content, err := http.GetContent(url)
	if err != nil {
		return nil, err
	}

	var tagStrings = strings.Join(tags, ",")
	newPage := model.Page{URL: url, Title: *title, Tags: tagStrings, Content: *content}
	err = repository.AddPage(newPage)
	if err != nil {
		return nil, err
	}
	return &newPage, err
}
