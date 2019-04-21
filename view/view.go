package view

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	page "github.com/kcwebapply/bm/page"
	util "github.com/kcwebapply/bm/util"
)

var (
	terminalWidth       int
	idPadding           = 3
	titlePadding        = 35
	urlPadding          int
	tagPadding          int
	minimumTerminalSize = 68
)

const (
	idColumnSize    = 3
	titleColumnSize = 30
	tagColumnSize   = 30
	idColor         = "\x1b[1m\x1b[38;5;181mid\x1b[0m"
	titleColor      = "\x1b[1m\x1b[38;5;112mtitle\x1b[0m"
	urlColor        = "\x1b[1m\x1b[38;5;133murl\x1b[0m"
	tagColor        = "\x1b[1m\x1b[38;5;216mtag\x1b[0m"
)

// this size is id-column size + title-columnt size + tagPadding.

func init() {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}

	terminalWidth = int(ws.Col)

	if terminalWidth > 165 {
		urlPadding = titlePadding + 100
	} else {
		urlPadding = terminalWidth - tagColumnSize
	}
	tagPadding = urlPadding + tagColumnSize
}

func printHeader() {
	nonPrintedCaracterSize := 0
	echo := "|"
	echo += idColor
	nonPrintedCaracterSize += 19
	echo = spacePadding(echo, "id", idPadding+nonPrintedCaracterSize)
	echo += "|"
	echo += titleColor
	nonPrintedCaracterSize += 19
	echo = spacePadding(echo, "title", titlePadding+nonPrintedCaracterSize)
	echo += "|"
	echo += urlColor
	nonPrintedCaracterSize += 19
	echo = spacePadding(echo, "", urlPadding+nonPrintedCaracterSize)
	echo += "|"
	echo += tagColor
	nonPrintedCaracterSize += 17
	echo = spacePadding(echo, "", tagPadding+nonPrintedCaracterSize)
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

// PrintTags printing Tags
func PrintTags(tagCounter map[string]int) {
	tags := make([]string, 0, len(tagCounter))
	for tag := range tagCounter {
		tags = append(tags, tag)
	}
	sort.Strings(tags) //sort by key
	var echo = ""
	var index = 0
	for _, tag := range tags {
		index++
		echo += fmt.Sprintf("%s:(%d), ", tag, tagCounter[tag])
		if index%5 == 0 {
			echo += "\n"
		}
	}
	fmt.Print(echo + "\n")
}

func printPage(data page.Page) {
	idString := strconv.Itoa(data.ID)
	echo := idString
	if data.ID < 10 {
		echo += "  "
	} else if data.ID < 100 {
		echo += " "
	}

	echo = spacePadding(echo, idString, idPadding)
	echo += "|"
	echo += data.Title
	echo = spacePadding(echo, data.Title, titlePadding)
	echo += "|"
	echo += shortURL(data.URL)
	echo = spacePadding(echo, shortURL(data.URL), urlPadding)
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
		space++
	}

	if space < 0 {
		sizingText := text[0:num]
		return sizingText
	}
	spaces := strings.Repeat(" ", space)
	return text + spaces
}

func tagView(tags []string) string {
	tagPrinter := strings.Join(tags, ",")
	return tagPrinter
}

func shortURL(url string) string {
	var urlReductedSize = urlPadding - titlePadding
	if len(url) >= urlReductedSize {
		shortURL := url[0 : urlReductedSize-5]
		shortURL += "..."
		return shortURL
	}
	return url
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
