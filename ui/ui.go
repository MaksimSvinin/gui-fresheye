package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/widget"
)

type UI struct {
	window fyne.Window
}

func NewUI(t []widget.RichTextSegment) *UI {
	a := app.New()
	w := a.NewWindow("Hello")

	// hello := widget.NewLabel("text")

	text := widget.NewRichText(t...)

	w.SetContent(text)

	return &UI{
		window: w,
	}
}

func (u *UI) Run() {
	u.window.ShowAndRun()
}
