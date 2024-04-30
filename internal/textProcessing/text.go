package textprocessing

import (
	"github.com/MaksimSvinin/gui-fresheye/internal/entity"
	"gitlab.com/opennota/fresheye"
)

func ProcessText(
	text string,
	sensitivityThreshold, contextSize, worldCount int,
	excludeProperNames bool,
) []entity.WorldInfo {
	it := fresheye.NewSimpleWordIterator(text)
	bad := fresheye.Check(it,
		fresheye.ContextSize(contextSize),
		fresheye.SensitivityThreshold(sensitivityThreshold),
		fresheye.ExcludeProperNames(excludeProperNames),
	)
	descriptors := fresheye.WordDescriptors(bad)

	promOut := make(map[string]entity.WorldInfo, 50)

	for _, d := range descriptors {
		from, to := d.Start(), d.End()
		world := text[from:to]

		if _, ok := promOut[world]; !ok {
			promOut[world] = entity.WorldInfo{
				Indexes: make([]entity.WorldIndex, 0, 10),
				World:   world,
			}
		}

		promOut[world] = entity.WorldInfo{
			Indexes: append(promOut[world].Indexes, entity.WorldIndex{
				From: from,
				To:   to,
			}),
			World: world,
		}
	}

	return MostRepeatedWords(promOut, worldCount)
}
