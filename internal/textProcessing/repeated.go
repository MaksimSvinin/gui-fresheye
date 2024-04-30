package textprocessing

import (
	"github.com/MaksimSvinin/gui-fresheye/internal/entity"
	"github.com/MaksimSvinin/gui-fresheye/internal/utils"
	"github.com/muesli/gamut"
)

func MostRepeatedWords(in map[string]entity.WorldInfo, count int) []entity.WorldInfo {
	countMap := make(map[int]map[string]entity.WorldInfo, 10)

	for _, worldInfo := range in {
		if _, ok := countMap[len(worldInfo.Indexes)]; !ok {
			countMap[len(worldInfo.Indexes)] = make(map[string]entity.WorldInfo)
		}

		countMap[len(worldInfo.Indexes)][worldInfo.World] = worldInfo
	}

	countArray := make([]int, 0, 10)
	for c := range countMap {
		countArray = append(countArray, c)
	}
	utils.Rsshell2(countArray)

	out := make([]entity.WorldInfo, 0, count)
	colors := gamut.Blends(gamut.Hex("#FF0000"), gamut.Hex("#00FF00"), count+1)

	i := 0
	for _, c := range countArray {
		for _, w := range countMap[c] {
			w.Count = c
			w.Color = colors[i]
			out = append(out, w)

			if i >= count {
				return out
			}

			i++
		}
	}

	return out
}
