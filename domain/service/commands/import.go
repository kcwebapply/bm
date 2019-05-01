package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/infrastructure/http"
)

// Import import chrome bookmark exports data on page-db.
func Import(c *cli.Context) {

	var bookmarkFilePath = c.Args().Get(0)

	// path validation
	if bookmarkFilePath == "" {
		log.Fatal("argument error. please input bookmark file on argument 1.")
	}

	imports(bookmarkFilePath)
}

func imports(bookmarkFilePath string) {

	var urlList = []string{}
	f, e := os.Open(bookmarkFilePath)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	// goquery
	var doc *goquery.Document
	if doc, e = goquery.NewDocumentFromReader(f); e != nil {
		log.Fatal(e)
	}

	//get bookmark urllist from bookmark file.
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urlList = append(urlList, url)
	})

	//page save loop
	for _, url := range urlList {
		title, content, err := http.GetContent(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		newPage := model.Page{URL: url, Title: *title, Content: *content}
		err = repository.AddPage(newPage)
		if err != nil {
			fmt.Println("err!", err)
			continue
		}
		fmt.Printf("page \"%s\" saved!\n", *title)
		time.Sleep(1 * time.Second)
	}

}
