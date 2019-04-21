package http

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// GetContent get content from internet.
func GetContent(u string) (*string, error) {
	baseURL, urlParseError := url.Parse(u)
	if urlParseError != nil {
		err := errors.New("url parse error! input url invalid")
		return nil, err
	}

	resp, httpGetError := http.Get(baseURL.String())
	if httpGetError != nil {
		err := errors.New("http get request error!")
		return nil, err
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	html, htmlGetError := doc.Html()
	if htmlGetError != nil {
		err := errors.New("content read error! contents are invalid")
		return nil, err
	}

	return &html, nil
}
