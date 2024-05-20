package ui

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/MaksimSvinin/gui-fresheye/internal/entity"
)

func showCheckWorlds(
	checkWorlds map[int]binding.Bool,
	worlds []entity.WorldInfo,
	inTextArea *widget.Entry,
	outTextArea *widget.RichText,
) {
	outTextArea.Segments = make([]widget.RichTextSegment, 0, 50)

	checkIndexMap := make(map[int]entity.WorldEndInfo)

	j := 0
	for i := range checkWorlds {
		b, err := checkWorlds[i].Get()
		if err != nil {
			continue
		}
		if b {
			w := worlds[i]

			for i := range w.Indexes {
				checkIndexMap[w.Indexes[i].From] = entity.WorldEndInfo{
					To:    w.Indexes[i].To,
					World: w.World,
					Color: w.Color,
				}
			}
			j++
		}
	}

	j = 0
	text := inTextArea.Text

	for i := range text {
		wi, ok := checkIndexMap[i]
		if !ok {
			continue
		}
		to := wi.To

		outTextArea.Segments = append(outTextArea.Segments,
			&widget.TextSegment{
				Style: widget.RichTextStyleCodeBlock,
				Text:  text[j:i],
			},
			&CustomSegment{
				Style: widget.RichTextStyleCodeBlock,
				Text:  text[i:to],
				Color: wi.Color,
			})
		j = to
	}

	outTextArea.Segments = append(outTextArea.Segments, &widget.TextSegment{
		Text:  text[j:],
		Style: widget.RichTextStyleCodeInline,
	})

	outTextArea.Refresh()
}
