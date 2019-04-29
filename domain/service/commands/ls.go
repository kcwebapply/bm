package commands

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// Ls returns bookmark list
func Ls(c *cli.Context) {
	pages := ls(c)
	view.PrintAllPage(pages)
}

func ls(c *cli.Context) []model.Page {
	var searchTitleParam = c.Args().Get(0)
	var tagSearchParam = c.String("t")
	var contentSearchParam = c.String("s")

	if searchTitleParam != "" {
		results, _ := repository.GetPagesByTitleWordGrep(searchTitleParam)
		return results
	}

	if tagSearchParam != "" {
		results, _ := repository.GetPagesByTag(tagSearchParam)
		return results
	}

	// content search case.
	if contentSearchParam != "" {
		word := contentSearchParam
		results, _ := repository.GetPagesByContentSearch(word)
		return results
	}

	pages, _ := repository.GetPages()
	return pages
}
