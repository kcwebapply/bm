package model

import (
	"strconv"
)

// Page type is struct of one page data.
type Page struct {
	ID      int
	URL     string
	Title   string
	Tags    string
	Content string
}

func (page *Page) String() string {
	return strconv.Itoa(page.ID) + " " + page.URL + " " + page.Title + " " + page.Tags + "\n"
}
