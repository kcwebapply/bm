package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// Ls returns bookmark list
func Ls(c *cli.Context) {

	pages, err := ls(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	view.PrintAllPage(*pages)
}

func ls(c *cli.Context) (*[]model.Page, error) {
	var searchTitleParam = c.Args().Get(0)
	var tagSearchParam = c.String("t")
	var contentSearchParam = c.String("s")

	if searchTitleParam != "" {
		results, err := repository.GetPagesByTitleWordGrep(searchTitleParam)
		if err != nil {
			return nil, err
		}
		return &results, err
	}

	if tagSearchParam != "" {
		results, err := repository.GetPagesByTag(tagSearchParam)
		if err != nil {
			return nil, err
		}
		return &results, err
	}

	// content search case.
	if contentSearchParam != "" {
		word := contentSearchParam
		results, err := repository.GetPagesByContentSearch(word)
		if err != nil {
			return nil, err
		}
		return &results, err
	}

	pages, err := repository.GetPages()
	return &pages, err
}
