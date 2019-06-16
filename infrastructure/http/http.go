package http

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/grokify/html-strip-tags-go"
)

// GetContent get content from internet.
func GetContent(u string) (*string, *string, error) {
	baseURL, urlParseError := url.Parse(u)
	if urlParseError != nil {
		return nil, nil, urlParseError
	}

	resp, httpGetError := http.Get(baseURL.String())
	if httpGetError != nil {
		return nil, nil, httpGetError
	}

	if resp.StatusCode != 200 {
		return nil, nil, fmt.Errorf("http request error ! %d", resp.StatusCode)
	}
	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	html, htmlGetError := doc.Html()
	if htmlGetError != nil {
		return nil, nil, htmlGetError
	}

	var title = doc.Find("title").Text()
	var content = removeHTMLTag(html)
	return &title, &content, nil
}

func removeHTMLTag(str string) string {
	src := strip.StripTags(str)
	src = strings.TrimSpace(src)
	return src
}
