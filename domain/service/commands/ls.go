package commands

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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

	pages, err := searchPages(c)

	// grep mode check
	var grepModeParam = c.Bool("g")
	if grepModeParam {
		pages, err = grep(*pages, err)
	}

	return pages, err
}

func searchPages(c *cli.Context) (*[]model.Page, error) {

	var searchTitleParam = c.Args().Get(0)
	var tagSearchParam = c.String("t")
	var contentSearchParam = c.String("s")

	// title search case.
	if searchTitleParam != "" {
		return searchByTitle(searchTitleParam)
	}

	// tag search case.
	if tagSearchParam != "" {
		return searchByTag(tagSearchParam)
	}

	// content search case.
	if contentSearchParam != "" {
		return searchByContent(contentSearchParam)
	}

	pages, err := repository.GetPages()
	if err != nil {
		log.Fatal(err)
	}

	return &pages, err
}

func searchByTitle(searchTitleParam string) (*[]model.Page, error) {

	results, err := repository.GetPagesByTitleWordGrep(searchTitleParam)
	if err != nil {
		return nil, err
	}

	return &results, err
}

func searchByTag(searchTagParam string) (*[]model.Page, error) {

	results, err := repository.GetPagesByTag(searchTagParam)
	if err != nil {
		return nil, err
	}

	return &results, err
}

func searchByContent(contentSearchParam string) (*[]model.Page, error) {

	results, err := repository.GetPagesByContentSearch(contentSearchParam)
	if err != nil {
		return nil, err
	}

	return &results, err
}

func grep(pages []model.Page, err error) (*[]model.Page, error) {

	view.PrintAllPage(pages)

	fmt.Println("(title grep) >>>")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		err = errors.New("get input error")
	}
	grepInput := scanner.Text()

	grepPages := []model.Page{}
	for _, page := range pages {
		if strings.Contains(page.Title, grepInput) {
			grepPages = append(grepPages, page)
		}
	}

	return &grepPages, err
}
