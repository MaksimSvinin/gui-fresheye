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
	closeLocCount int,
	closeWorldFlag bool,
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
					To:       w.Indexes[i].To,
					World:    w.World,
					Color:    w.Color,
					CloseLoc: false,
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
			})

		customText := text[i:to]
		closeLoc := wi.CloseLoc

		for k := to; k < to+closeLocCount; k++ {
			wik, okk := checkIndexMap[k]
			if !okk || wik.World != wi.World {
				continue
			}
			closeLoc = true
			checkIndexMap[k] = entity.WorldEndInfo{
				To:       wik.To,
				World:    wik.World,
				Color:    wik.Color,
				CloseLoc: true,
			}
		}
		appendWorld(closeLoc, closeWorldFlag, wi, customText, outTextArea)

		j = to
	}

	outTextArea.Segments = append(outTextArea.Segments, &widget.TextSegment{
		Text:  text[j:],
		Style: widget.RichTextStyleCodeInline,
	})

	outTextArea.Refresh()
}

func appendWorld(
	closeLoc, closeWorldFlag bool,
	wi entity.WorldEndInfo,
	customText string,
	outTextArea *widget.RichText,
) {
	if closeLoc {
		outTextArea.Segments = append(outTextArea.Segments, &CustomSegment{
			Style: widget.RichTextStyleHeading,
			Text:  customText,
			Color: wi.Color,
		})
	} else {
		if closeWorldFlag {
			outTextArea.Segments = append(outTextArea.Segments, &widget.TextSegment{
				Style: widget.RichTextStyleInline,
				Text:  customText,
			})
		} else {
			outTextArea.Segments = append(outTextArea.Segments, &CustomSegment{
				Style: widget.RichTextStyleInline,
				Text:  customText,
				Color: wi.Color,
			})
		}
	}
}
