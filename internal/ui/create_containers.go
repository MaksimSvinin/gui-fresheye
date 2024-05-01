package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func createBorder(
	sensitivityThresholdEntry, contextSizeEntry, worldCountEntry *widget.Entry,
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
		excludeProperNamesCheck,
		win1251check,
	)
}

func createEntry() (*widget.Entry, *widget.Entry, *widget.Entry) {
	sensitivityThresholdEntry := widget.NewEntry()
	contextSizeEntry := widget.NewEntry()
	worldCountEntry := widget.NewEntry()

	sensitivityThresholdEntry.Text = "100"
	contextSizeEntry.Text = "100"
	worldCountEntry.Text = "20"

	return sensitivityThresholdEntry, contextSizeEntry, worldCountEntry
}
