package ui

import (
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type customLayout struct {
	CellSize fyne.Size
	colCount int
	rowCount int
}

func NewCustomLayout(widht float32) fyne.Layout {
	return &customLayout{fyne.NewSize(widht, 0), 1, 1}
}

func (g *customLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	padding := theme.Padding()
	g.colCount = 1
	g.rowCount = 0

	g.CellSize.Height = size.Height

	if size.Width > g.CellSize.Width {
		g.colCount = int(math.Floor(float64(size.Width+padding) / float64(g.CellSize.Width+padding)))
	}

	i, x, y := 0, float32(0), float32(0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if i%g.colCount == 0 {
			g.rowCount++
		}

		child.Move(fyne.NewPos(x, y))
		child.Resize(g.CellSize)

		if (i+1)%g.colCount == 0 {
			x = 0
			y += g.CellSize.Height + padding
		} else {
			x += g.CellSize.Width + padding
		}
		i++
	}
}

func (g *customLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	rows := g.rowCount
	if rows < 1 {
		rows = 1
	}
	return fyne.NewSize(g.CellSize.Width,
		(g.CellSize.Height*float32(rows))+(float32(rows-1)*theme.Padding()))
}
