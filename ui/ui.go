package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/widget"
)

type Ui struct {
	window fyne.Window
}

func NewUi(t []widget.RichTextSegment) *Ui {
	a := app.New()
	w := a.NewWindow("Hello")

	//hello := widget.NewLabel("text")

	text := widget.NewRichText(t...)

	w.SetContent(text)

	return &Ui{
		window: w,
	}
}

func (u *Ui) Run() {
	u.window.ShowAndRun()
}
