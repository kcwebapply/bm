package commands

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
	page "github.com/kcwebapply/bm/page"
	util "github.com/kcwebapply/bm/util"
	view "github.com/kcwebapply/bm/view"
	homedir "github.com/mitchellh/go-homedir"
	open "github.com/skratchdot/open-golang/open"
)

var fileName = ""

var maxTextSize = 60

func init() {
	filePath, _ := homedir.Dir()
	fileName = fmt.Sprintf("%s/%s", filePath, "bm.txt")
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

	newData := savePage(url, title, tagList)

	view.PrintSavePage(newData)
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
	writer := getFileCleanWriter()
	defer writer.Flush()
	pageSize := len(allPages)
	for _, page := range allPages {
		writer.Write(([]byte)(page.String()))
	}

	var lastCOUNTER int
	if pageSize > 0 {
		lastCOUNTER = allPages[pageSize-1].ID + 1
	} else {
		lastCOUNTER = 1
	}
	newData := page.Page{ID: lastCOUNTER, URL: url, Title: title, Tags: tagList}
	writer.Write(([]byte)(newData.String()))
	return newData
}

func deletePage(id string) page.Page {
	allPages := readLines()
	writer := getFileCleanWriter()
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

func getFileCleanWriter() *bufio.Writer {
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
