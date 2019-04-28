package repository

import (
	"sort"

	"github.com/kcwebapply/bm/page"
)

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
