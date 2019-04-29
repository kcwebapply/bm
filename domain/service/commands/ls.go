package commands

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/domain/model"
	"github.com/kcwebapply/bm/domain/repository"
	"github.com/kcwebapply/bm/view"
)

// Ls returns bookmark list
func Ls(c *cli.Context) {
	search := c.Args().Get(0)
	allPages := []model.Page{}

	if search != "" {
		allPages, _ = repository.GetPagesByTitleWordGrep(search)
	} else {
		allPages, _ = repository.GetPages()
	}

	if c.String("t") != "" {
		searchPages := []model.Page{}
		searchTag := c.String("t")
		for _, page := range allPages {
			tags := strings.Split(page.Tags, ",")
			for _, tag := range tags {
				if searchTag == tag {
					searchPages = append(searchPages, page)
					break
				}
			}
		}
		allPages = searchPages
	}

	if c.String("s") != "" {
		word := c.String("s")
		allPages, _ = repository.GetPagesByTitleWordGrep(word)
	}

	view.PrintAllPage(allPages)
}

func searchPageContent(word string, allPages []model.Page) []model.Page {
	var results = []model.Page{}
	/*for _, page := range allPages {
		ID := page.ID
		pageContentFile := contentPath + "/" + strconv.Itoa(ID) + ".txt"
		command := fmt.Sprintf("cat %s | grep %s", pageContentFile, word)
		_, err := exec.Command("sh", "-c", command).Output()
		if err != nil {
			continue
		}
		results = append(results, page)
	}*/
	return results
}
