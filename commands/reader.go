package commands

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	page "github.com/kcwebapply/bm/page"
)

func readPages() []page.Page {
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

func readPagesByTitleWordGrep(word string) []page.Page {
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
