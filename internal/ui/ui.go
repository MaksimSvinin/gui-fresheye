package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/MaksimSvinin/gui-fresheye/internal/entity"
)

type UI struct {
	window fyne.Window
}

func NewUI() *UI {
	a := app.New()
	w := a.NewWindow("gui fresheye")
	w.Resize(fyne.NewSize(500, 500))

	sensitivityThresholdEntry, contextSizeEntry, worldCountEntry := createEntry()
	excludeProperNames := false
	excludeProperNamesCheck := widget.NewCheck("исключить имена собственные", func(b bool) {
		excludeProperNames = b
	})

	worlds := make([]entity.WorldInfo, 0, 10)
	checkWorlds := make(map[int]binding.Bool)

	inTextArea := widget.NewEntry()
	outTextArea := widget.NewRichText()
	errorArea := widget.NewLabel("")

	worldsList := widget.NewList(
		func() int {
			return len(worlds)
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("", func(b bool) {})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Check).SetText(worlds[lii].World + " " + strconv.Itoa(worlds[lii].Count))
			if _, ok := checkWorlds[lii]; !ok {
				checkWorlds[lii] = binding.NewBool()
			}
			co.(*widget.Check).Bind(checkWorlds[lii])
		},
	)

	fileMenu := fyne.NewMenu("file",
		fyne.NewMenuItem("Open file", func() {
			dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
				updateError(err, errorArea)
				if err != nil {
					return
				}
				text, err := readFile(uc)
				updateError(err, errorArea)
				if err != nil {
					return
				}
				inTextArea.SetText(text)
			}, w)
		}),
	)
	w.SetMainMenu(fyne.NewMainMenu(fileMenu))

	border := createBorder(sensitivityThresholdEntry, contextSizeEntry, worldCountEntry, excludeProperNamesCheck)

	analyzeButton := widget.NewButton("analyze", func() {
		var err error
		worlds, err = analyze(
			inTextArea.Text,
			sensitivityThresholdEntry,
			contextSizeEntry,
			worldCountEntry,
			excludeProperNames,
		)
		updateError(err, errorArea)
		if err != nil {
			return
		}
		worldsList.Refresh()
	})

	showButton := widget.NewButton("show", func() {
		showCheckWorlds(checkWorlds, worlds, inTextArea, outTextArea)
	})

	showAnalyzeButtons := container.New(
		layout.NewGridLayoutWithRows(2),
		analyzeButton,
		showButton,
	)

	w.SetContent(
		container.New(
			layout.NewBorderLayout(border, errorArea, nil, nil),
			border,
			errorArea,
			container.New(layout.NewGridLayoutWithColumns(2), container.New(layout.NewGridLayoutWithRows(2),
				container.NewScroll(inTextArea), container.NewScroll(outTextArea)),
				container.New(layout.NewBorderLayout(showAnalyzeButtons, nil, nil, nil), showAnalyzeButtons, worldsList),
			),
		),
	)

	return &UI{window: w}
}

func (u *UI) Run() {
	u.window.ShowAndRun()
}
