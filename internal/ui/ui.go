package ui

import (
	"embed"
	"runtime"
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

const infoMsg = `Программа для поиска близко расположенных одинаковых слов (повторов) 
и частотного анализа употребления слов в русском тексте.`

func NewUI(f embed.FS) *UI {
	a := app.New()
	w := a.NewWindow("gui fresheye")
	w.Resize(fyne.NewSize(700, 700))

	sensitivityThresholdEntry, contextSizeEntry, worldCountEntry, closeLocCountEntry := createEntry()
	excludeProperNames := false
	excludeProperNamesCheck := widget.NewCheck("Исключить имена собственные", func(b bool) {
		excludeProperNames = b
	})
	win1251 := false
	win1251check := widget.NewCheck("win1251 кодировка", func(b bool) {
		win1251 = b
	})
	if runtime.GOOS == "windows" {
		win1251check.SetChecked(true)
	}

	worlds := make([]entity.WorldInfo, 0, 10)
	checkWorlds := make(map[int]binding.Bool)

	inTextArea := &widget.Entry{MultiLine: true, Wrapping: fyne.TextWrapBreak}
	outTextArea := &widget.RichText{Wrapping: fyne.TextWrapBreak}
	errorArea := widget.NewLabel("")
	closeWorldFlag := false

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

	setMenu(errorArea, &win1251, inTextArea, w, a, f)
	border := createBorder(sensitivityThresholdEntry, contextSizeEntry, worldCountEntry, closeLocCountEntry,
		excludeProperNamesCheck, win1251check)

	analyzeButton := widget.NewButton("Анализ", func() {
		checkWorlds = make(map[int]binding.Bool)
		var err error
		worlds, err = analyze(inTextArea.Text, sensitivityThresholdEntry, contextSizeEntry,
			worldCountEntry, excludeProperNames)
		updateError(err, errorArea)
		if err != nil {
			return
		}
		worldsList.Refresh()
	})

	showButton := widget.NewButton("Показать выделенные слова", func() {
		closeLocCount, err := strconv.Atoi(closeLocCountEntry.Text)
		updateError(err, errorArea)
		if err != nil {
			return
		}
		showCheckWorlds(checkWorlds, worlds, inTextArea, outTextArea, closeLocCount, closeWorldFlag)
	})

	closeWorld := widget.NewCheck("Показывать только близкие слова", func(b bool) { closeWorldFlag = b })
	selectAll := widget.NewCheck("Выделить всё", func(b bool) { selectAll(b, checkWorlds, errorArea) })

	mainContainer := createMainContainer(analyzeButton, showButton, closeWorld, selectAll, worldsList)
	w.SetContent(
		container.New(
			layout.NewBorderLayout(border, errorArea, nil, mainContainer),
			border,
			errorArea,
			container.New(layout.NewGridLayoutWithColumns(2),
				container.NewScroll(inTextArea), container.NewScroll(outTextArea)),
			mainContainer,
		),
	)

	return &UI{window: w}
}

func (u *UI) Run() {
	u.window.ShowAndRun()
}

func setMenu(
	errorArea *widget.Label,
	win1251 *bool,
	inTextArea *widget.Entry,
	w fyne.Window,
	a fyne.App,
	f embed.FS,
) {
	fileMenu := fyne.NewMenu("Фаил",
		fyne.NewMenuItem("Открыть фаил", func() {
			dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
				updateError(err, errorArea)
				if err != nil {
					return
				}
				text, err := readFile(uc, *win1251)
				updateError(err, errorArea)
				if err != nil {
					return
				}
				inTextArea.SetText(text)
			}, w)
		}),
	)
	infoMenu := fyne.NewMenu("Информация",
		fyne.NewMenuItem("О программе", func() {
			dialog.ShowForm("О программе", "", "", []*widget.FormItem{
				{
					Text:   "name:",
					Widget: widget.NewLabel(a.Metadata().Name),
				},
				{
					Text:   "info",
					Widget: widget.NewLabel(infoMsg),
				},
				{
					Text:   "version:",
					Widget: widget.NewLabel(a.Metadata().Version),
				},
			}, func(b bool) {}, w)
		}),
		fyne.NewMenuItem("Справка", func() {
			dialog.ShowForm("Справка", "", "", readReadmeFile(f), func(b bool) {}, w)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(fileMenu, infoMenu))
}

func readReadmeFile(f embed.FS) []*widget.FormItem {
	readmeFile, _ := f.ReadFile("README.md")

	return []*widget.FormItem{
		{
			Text:   "",
			Widget: widget.NewRichTextFromMarkdown(string(readmeFile)),
		},
	}
}

func selectAll(
	b bool,
	checkWorlds map[int]binding.Bool,
	errorArea *widget.Label,
) {
	for i := range checkWorlds {
		if err := checkWorlds[i].Set(b); err != nil {
			updateError(err, errorArea)
		}
	}
}
