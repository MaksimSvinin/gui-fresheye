package ui

import (
	"strconv"

	"fyne.io/fyne/v2/widget"
)

func readEntry(sensitivityThreshold, contextSize, worldCount *widget.Entry) (int, int, int, error) {
	intSensitivityThreshold, err := strconv.Atoi(sensitivityThreshold.Text)
	if err != nil {
		return 0, 0, 0, err
	}

	intContextSize, err := strconv.Atoi(contextSize.Text)
	if err != nil {
		return 0, 0, 0, err
	}

	intWorldCount, err := strconv.Atoi(worldCount.Text)
	if err != nil {
		return 0, 0, 0, err
	}
	return intSensitivityThreshold, intContextSize, intWorldCount, nil
}
