package commands

import (
	"github.com/codegangsta/cli"
	page "github.com/kcwebapply/bm/page"
	repository "github.com/kcwebapply/bm/repository"
	view "github.com/kcwebapply/bm/view"
)

func Ls(c *cli.Context) {
	search := c.Args().Get(0)
	allPages := []page.Page{}

	if search != "" {
		allPages = repository.GetPagesByTitleWordGrep(search)
	} else {
		allPages = repository.GetPages()
	}

	if c.String("t") != "" {
		searchPages := []page.Page{}
		searchTag := c.String("t")
		for _, page := range allPages {
			for _, tag := range page.Tags {
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
		allPages = searchPageContent(word, allPages)
	}

	view.PrintAllPage(allPages)
}

func searchPageContent(word string, allPages []page.Page) []page.Page {
	var results = []page.Page{}
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
