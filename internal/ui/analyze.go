package ui

import (
	"github.com/MaksimSvinin/gui-fresheye/internal/entity"
	textprocessing "github.com/MaksimSvinin/gui-fresheye/internal/textProcessing"

	"fyne.io/fyne/v2/widget"
)

func analyze(
	text string,
	sensitivityThresholdEntry, contextSizeEntry, worldCountEntry *widget.Entry,
	excludeProperNames bool,
) ([]entity.WorldInfo, error) {
	sensitivityThreshold, contextSize, worldCount, err := readEntry(
		sensitivityThresholdEntry,
		contextSizeEntry,
		worldCountEntry,
	)
	if err != nil {
		return nil, err
	}

	return textprocessing.ProcessText(
		text,
		sensitivityThreshold,
		contextSize,
		worldCount,
		excludeProperNames,
	), nil
}
