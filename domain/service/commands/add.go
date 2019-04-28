package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/util"
	"github.com/kcwebapply/bm/view"
)

// Add saves pagedata
func Add(c *cli.Context) {

	url := c.Args().Get(0)
	tagList := []string{}

	title, _ := util.GetTerminalInput("title")

	tagsInput, _ := util.GetTerminalInput("tags (input few tags by ',')")
	tags := strings.Split(tagsInput, ",")
	tagSize := len(tags)
	if tagSize > 0 && tagSize <= 3 {
		tagList = tags
	} else if tagSize > 3 {
		fmt.Println("tag size error. you can't put more than 3 tags on your bookmark.")
		os.Exit(0)
	}

	newPage := add(url, title, tagList)
	// save http content to ${home}/${ID}.txt
	saveHTTPContent(newPage.ID, newPage.URL)

	view.PrintRm(newPage)
}

func add(url string, title string, tagList []string) model.Page {
	allPages := repository.GetPages()
	pageSize := len(allPages)
	var newID = 1
	if pageSize > 0 {
		newID = allPages[pageSize-1].ID + 1
	}

	newPage := model.Page{ID: newID, URL: url, Title: title, Tags: tagList}
	repository.AddPage(newPage)
	return newPage
}

func saveHTTPContent(id int, url string) {
	// get page content from www.
	/*content, err := http.GetContent(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}*/
	/*ontentFileName := contentPath + "/" + strconv.Itoa(id) + ".txt"
	writer := getFileCleanWriter(contentFileName)
	defer writer.Flush()
	writer.Write(([]byte)(*content))*/
}
