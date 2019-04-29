package commands

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// Tags returns all tags on bookmarks
func Tags(c *cli.Context) {
	allPages, _ := repository.GetPages()
	tagCounter := make(map[string]int)
	for _, page := range allPages {
		//if _, ok := idSets[data.Id]; !ok {
		tags := strings.Split(page.Tags, ",")
		for _, tag := range tags {
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
