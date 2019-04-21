package page

import (
	"errors"
	"strconv"
	"strings"
)

// Page type is struct of one page data.
type Page struct {
	ID    int
	URL   string
	Title string
	Tags  []string
}

func (page *Page) String() string {
	tagStrings := strings.Join(page.Tags, " ")
	return strconv.Itoa(page.ID) + " " + page.URL + " " + page.Title + " " + tagStrings + "\n"
}

// ConvertToPage is function for converting string page expression to Page struct.
func ConvertToPage(page string) (Page, error) {
	var data Page
	var err error
	datas := strings.Split(page, " ")
	url := datas[1]
	title := datas[2]
	tags := []string{}
	if len(datas) < 3 {
		err = errors.New("data parse error")
	} else {
		id, _ := strconv.Atoi(datas[0])
		if id == 0 {
			err = errors.New("malforlmed data error")
		}
		if len(datas) > 3 {
			tags = datas[3:]
		}
		data = Page{ID: id, URL: url, Title: title, Tags: tags}
	}
	return data, err
}
