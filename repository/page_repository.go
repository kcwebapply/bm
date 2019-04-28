package repository

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	page "github.com/kcwebapply/bm/page"
)

//GetPages returns all page entities user saved.
func GetPages() []page.Page {
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

// GetPagesByTitleWordGrep retunrs page grepped title by input word.
func GetPagesByTitleWordGrep(word string) []page.Page {
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

// AddPage saved bookrmark user input.
func AddPage(newPage page.Page) error {
	allPages := GetPages()
	fileWriter := getFileCleanWriter()
	defer fileWriter.Flush()
	for _, page := range allPages {
		fileWriter.Write(([]byte)(page.String()))
	}

	fileWriter.Write(([]byte)(newPage.String()))
	return nil
}

// RemovePage remove bookmark
func RemovePage(id string) (page.Page, error) {
	allPages := GetPages()
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

	return deletePage, nil
}

func getFileCleanWriter() *bufio.Writer {
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	return writer
}
