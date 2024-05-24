package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createBorder(
	sensitivityThresholdEntry, contextSizeEntry, worldCountEntry, closeLocCountEntry *widget.Entry,
	excludeProperNamesCheck, win1251check *widget.Check,
) *fyne.Container {
	return container.New(
		layout.NewHBoxLayout(),
		sensitivityThresholdEntry,
		widget.NewLabel("Порог чувствительности"),
		contextSizeEntry,
		widget.NewLabel("Размер контекста"),
		worldCountEntry,
		widget.NewLabel("Количество выводимых слов"),
		closeLocCountEntry,
		widget.NewLabel("Растояние между парами слов в символах"),
		excludeProperNamesCheck,
		win1251check,
	)
}

func createEntry() (*widget.Entry, *widget.Entry, *widget.Entry, *widget.Entry) {
	sensitivityThresholdEntry := widget.NewEntry()
	contextSizeEntry := widget.NewEntry()
	worldCountEntry := widget.NewEntry()
	closeLocCountEntry := widget.NewEntry()

	sensitivityThresholdEntry.Text = "100"
	contextSizeEntry.Text = "100"
	worldCountEntry.Text = "20"
	closeLocCountEntry.Text = "50"

	return sensitivityThresholdEntry, contextSizeEntry, worldCountEntry, closeLocCountEntry
}

func createMainContainer(
	analyzeButton, showButton *widget.Button,
	closeWorld, selectAll *widget.Check,
	worldsList *widget.List,
) *fyne.Container {
	showAnalyzeButtons := container.New(
		layout.NewGridLayoutWithRows(4),
		analyzeButton,
		showButton,
		closeWorld,
		selectAll,
	)

	return container.New(
		NewCustomLayout(300),
		container.New(layout.NewBorderLayout(showAnalyzeButtons, nil, nil, nil), showAnalyzeButtons, worldsList),
	)
}
