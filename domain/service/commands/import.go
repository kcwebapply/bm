package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/infrastructure/http"
	"github.com/kcwebapply/bm/util"
)

// Import import chrome bookmark exports data on page-db.
func Import(c *cli.Context) {

	var bookmarkFilePath = c.Args().Get(0)

	imports(bookmarkFilePath)
}

func imports(bookmarkFilePath string) {

	f, e := os.Open(bookmarkFilePath)
	if e != nil {
		util.LoggingError(e.Error())
	}
	defer f.Close()
	// goquery
	var doc *goquery.Document
	if doc, e = goquery.NewDocumentFromReader(f); e != nil {
		util.LoggingError(e.Error())
	}

	//get bookmark urllist from bookmark file.
	var urlList = []string{}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urlList = append(urlList, url)
	})

	//page save loop
	for _, url := range urlList {

		title, content, err := http.GetContent(url)
		if err != nil {
			fmt.Printf("http error getting : %s \n ", url)
			continue
		}

		newPage := model.Page{URL: url, Title: *title, Content: *content}
		err = repository.AddPage(newPage)
		if err != nil {
			fmt.Printf("saving page error : %s \n ", err.Error())
			continue
		}
		fmt.Printf("page \"%s\" saved!\n", *title)

		time.Sleep(1 * time.Second)
	}

}
