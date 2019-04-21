package commands

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	http "github.com/kcwebapply/bm/http"
	page "github.com/kcwebapply/bm/page"
	provider "github.com/kcwebapply/bm/provider"
	util "github.com/kcwebapply/bm/util"
	view "github.com/kcwebapply/bm/view"
	open "github.com/skratchdot/open-golang/open"
)

var (
	fileName    = ""
	maxTextSize = 60

	contentPath = ""
)

func init() {
	fileName = provider.FileName
	contentPath = provider.ContentPath
}

func GetAllPages(c *cli.Context) {
	search := c.Args().Get(0)
	allPages := []page.Page{}

	if search != "" {
		allPages = readLinesBySearch(search)
	} else {
		allPages = readLines()
	}

	if c.String("t") != "" {
		searchPages := []page.Page{}
		searchTag := c.String("t")
		for _, page := range allPages {
			for _, tag := range page.Tags {
				if searchTag == tag {
					searchPages = append(searchPages, page)
					break
				}
			}
		}
		allPages = searchPages
	}

	view.PrintAllPage(allPages)
}

// GetTags returns all tags on bookmarks
func GetTags(c *cli.Context) {
	allPages := readLines()
	tagCounter := make(map[string]int)
	for _, page := range allPages {
		//if _, ok := idSets[data.Id]; !ok {
		for _, tag := range page.Tags {
			if tag == "" {
				continue
			}
			if _, contain := tagCounter[tag]; !contain {
				tagCounter[tag] = 1
			} else {
				tagCounter[tag]++
			}
		}
	}
	view.PrintTags(tagCounter)

}

// SavePage saves pagedata
func SavePage(c *cli.Context) {

	url := c.Args().Get(0)
	fmt.Println("url:", url)
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

	newPage := savePage(url, title, tagList)
	// save http content to ${home}/${ID}.txt
	saveHTTPContent(newPage.ID, newPage.URL)

	view.PrintSavePage(newPage)
}

// DeletePage delete pagedata
func DeletePage(c *cli.Context) {

	id := c.Args().Get(0)
	if id == "" {
		os.Exit(0)
	}
	page := deletePage(id)

	view.PrintSavePage(page)
}

// OpenPage open bookmark website
func OpenPage(c *cli.Context) {
	id := c.Args().Get(0)
	if id == "" {
		os.Exit(0)
	}
	openPage(id)

}

func savePage(url string, title string, tagList []string) page.Page {
	allPages := readLines()
	fileWriter := getFileCleanWriter(fileName)
	defer fileWriter.Flush()
	pageSize := len(allPages)
	for _, page := range allPages {
		fileWriter.Write(([]byte)(page.String()))
	}

	var newID int
	if pageSize > 0 {
		newID = allPages[pageSize-1].ID + 1
	} else {
		newID = 1
	}

	newPage := page.Page{ID: newID, URL: url, Title: title, Tags: tagList}
	fileWriter.Write(([]byte)(newPage.String()))
	return newPage
}

func saveHTTPContent(id int, url string) {
	// get page content from www.
	content, err := http.GetContent(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	contentFileName := contentPath + "/" + strconv.Itoa(id) + ".txt"
	writer := getFileCleanWriter(contentFileName)
	writer.Write(([]byte)(*content))
}

func deletePage(id string) page.Page {
	allPages := readLines()
	writer := getFileCleanWriter(fileName)
	defer writer.Flush()

	var deletePage page.Page

	for _, page := range allPages {
		if strconv.Itoa(page.ID) == id {
			deletePage = page
			continue
		}
		writer.Write(([]byte)(page.String()))
	}

	return deletePage
}

func openPage(id string) {
	pages := readLines()
	for _, page := range pages {
		if id == strconv.Itoa(page.ID) {
			open.Run(page.URL)
		}
	}
}

func readLines() []page.Page {
	var lines = []page.Page{}

	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data, err := page.ConvertToPage(scanner.Text())
		if err != nil {
			continue
		}
		lines = append(lines, data)
	}
	lines = sortAndDeleteDuplication(lines)
	return lines
}

func readLinesBySearch(word string) []page.Page {
	var lines = []page.Page{}

	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data, err := page.ConvertToPage(scanner.Text())

		if !strings.Contains(data.Title, word) {
			continue
		}
		if err != nil {
			continue
		}
		lines = append(lines, data)
	}
	lines = sortAndDeleteDuplication(lines)
	return lines
}

func getFileCleanWriter(fileName string) *bufio.Writer {
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	return writer
}

func sortAndDeleteDuplication(datas []page.Page) []page.Page {
	dataSets := []page.Page{}
	idSets := make(map[int]struct{})

	for _, data := range datas {
		if _, ok := idSets[data.ID]; !ok {
			idSets[data.ID] = struct{}{}
			dataSets = append(dataSets, data)
		}
	}

	sort.Slice(dataSets, func(i, j int) bool {
		return dataSets[i].ID < dataSets[j].ID
	})
	return dataSets
}
