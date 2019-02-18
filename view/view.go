package view

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	page "github.com/kcwebapply/bm/page"
	"github.com/kcwebapply/bm/util"
)

var idPadding = 3
var titlePadding = 35
var urlPadding = 135
var tagPadding = 165

func printHeader() {
	echo := ""
	echo += "|\x1b[1m\x1b[38;5;181mid\x1b[0m"
	echo = spacePadding(echo, "id", idPadding+19)
	echo += "|\x1b[1m\x1b[38;5;112mtitle\x1b[0m"
	echo = spacePadding(echo, "title", titlePadding+38)
	echo += "|\x1b[1m\x1b[38;5;133murl\x1b[0m"
	echo = spacePadding(echo, "", urlPadding+57)
	echo += "|\x1b[1m\x1b[38;5;216mtag\x1b[0m"
	echo = spacePadding(echo, "", tagPadding+74)
	echo += "|"
	line := strings.Repeat("-", len(echo)-76)
	fmt.Println(line)
	fmt.Println(echo)
	fmt.Println(line)
}

// PrintAllMemoPage is function of printing message when showing all page.
func PrintAllPage(datas []page.Page) {
	printHeader()
	for _, data := range datas {
		printPage(data)
	}
}

// PrintSavePage is function of printing message when saving page.
func PrintSavePage(data page.Page) {
	printHeader()
	printPage(data)
	fmt.Println("\x1b[1m\x1b[38;5;39mbookmark completed!\x1b[0m")
}

// PrintDeletePage is function of printing message when deleting page.
func PrintDeletePage(data page.Page) {
	printHeader()
	printPage(data)
	fmt.Println("\x1b[1m\x1b[38;5;39mbookmark deleted!\x1b[0m")
}

func PrintTags(tagCounter map[string]int) {
	tags := make([]string, 0, len(tagCounter))
	for tag := range tagCounter {
		tags = append(tags, tag)
	}
	sort.Strings(tags) //sort by key
	var echo = ""
	var index = 0
	for _, tag := range tags {
		index += 1
		echo += fmt.Sprintf("%s:(%d), ", tag, tagCounter[tag])
		if index%5 == 0 {
			echo += "\n"
		}
	}
	fmt.Print(echo + "\n")
}

func printPage(data page.Page) {
	idString := strconv.Itoa(data.Id)
	echo := idString
	if data.Id < 10 {
		echo += "  "
	} else if data.Id < 100 {
		echo += " "
	}

	echo = spacePadding(echo, idString, idPadding)
	echo += "|"
	echo += data.Title
	echo = spacePadding(echo, data.Title, titlePadding)
	echo += "|"
	echo += shortUrl(data.URL)
	echo = spacePadding(echo, shortUrl(data.URL), urlPadding)
	echo += "|"
	tagString := tagView(data.Tags)
	echo += tagString
	echo = spacePadding(echo, tagString, tagPadding)
	fmt.Println(echo)
}

func spacePadding(text string, content string, num int) string {
	textLength := util.TextCounter(text)
	space := num - textLength
	rep := regexp.MustCompile("^([a-zA-Z0-9])+$")

	if rep.MatchString(content) {
		space += 1
	}

	if space < 0 {
		sizingText := text[0:num]
		return sizingText
	} else {
		spaces := strings.Repeat(" ", space)
		return text + spaces
	}
}

func tagView(tags []string) string {
	tagPrinter := strings.Join(tags, ",")
	return tagPrinter
}

func shortUrl(url string) string {
	if len(url) >= 100 {
		shortURL := url[0:95]
		shortURL += "..."
		return shortURL
	}
	return url
}
