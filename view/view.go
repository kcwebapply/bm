package view

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/kcwebapply/bm/domain/model"
)

var (
	terminalWidth       int
	idPadding           = 3
	titlePadding        = 135
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
		titlePadding = 132
	} else {
		titlePadding = terminalWidth - (40 + tagColumnSize)
	}
	tagPadding = terminalWidth - 40
}

func printHeader() {
	nonPrintedCaracterSize := 0
	echo := ""
	echo += idColor
	nonPrintedCaracterSize += 19
	echo = spacePadding(echo, "id", idPadding+nonPrintedCaracterSize)
	echo += "|"
	echo += titleColor
	nonPrintedCaracterSize += 19
	echo = spacePadding(echo, "title", titlePadding+nonPrintedCaracterSize)
	echo = spacePadding(echo, "", terminalWidth-40+nonPrintedCaracterSize)
	echo += "|"
	echo += tagColor

	line := strings.Repeat("-", terminalWidth-10)
	//fmt.Println(line)
	fmt.Println(echo)
	fmt.Println(line)
}

// PrintAllPage is function of printing message when showing all page.
func PrintAllPage(datas []model.Page) {
	printHeader()
	for _, page := range datas {
		printPage(page)
	}
}

// PrintAdd is function of printing message when saving page.
func PrintAdd(data model.Page) {
	printHeader()
	printPage(data)
	fmt.Println("\x1b[1m\x1b[38;5;39mbookmark completed!\x1b[0m")
}

// PrintRm is function of printing message when deleting page.
func PrintRm(id string) {
	fmt.Printf("\x1b[1m\x1b[38;5;39mbookmark %s deleted!\x1b[0m\n", id)
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

func printPage(data model.Page) {
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
	echo = spacePadding(echo, shortTitle(data.Title), titlePadding)
	tagString := data.Tags
	echo = spacePadding(echo, "", terminalWidth-40)
	echo += "|"
	echo += tagString
	fmt.Println(echo)
}

func spacePadding(text string, content string, num int) string {
	textLength := textCounter(text)
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

func shortTitle(title string) string {
	var titleReductedSize = titlePadding - idPadding
	if len(title) >= titleReductedSize {
		shortTitle := title[0 : titleReductedSize-5]
		shortTitle += "..."
		return shortTitle
	}
	return title
}

func textCounter(text string) int {
	textCounter := 0
	befPos := 0
	for pos := range text {
		if pos-befPos == 3 {
			textCounter += 2 // to treat japanese character as 2byte.
			befPos = pos
		} else {
			textCounter++
			befPos = pos
		}
	}
	return textCounter
}

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
