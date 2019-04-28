package repository

import (
	"sort"

	"github.com/kcwebapply/bm/domain/model"
)

func sortAndDeleteDuplication(datas []model.Page) []model.Page {
	dataSets := []model.Page{}
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
