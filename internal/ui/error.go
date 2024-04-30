package ui

import "fyne.io/fyne/v2/widget"

func updateError(err error, errorArea *widget.Label) {
	if err == nil {
		errorArea.SetText("")
	} else {
		errorArea.SetText(err.Error())
	}
}
