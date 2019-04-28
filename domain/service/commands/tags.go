package commands

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// GetTags returns all tags on bookmarks
func GetTags(c *cli.Context) {
	allPages := repository.GetPages()
	tagCounter := make(map[string]int)
	for _, page := range allPages {
		//if _, ok := idSets[data.Id]; !ok {
		for _, tag := range page.Tags {
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
