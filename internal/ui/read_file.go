package ui

import (
	"errors"
	"io"
	"path"

	"fyne.io/fyne/v2"
)

func readFile(uc fyne.URIReadCloser) (string, error) {
	if uc == nil {
		return "", errors.New("not select file")
	}
	defer uc.Close()
	if err := validateExt(uc.URI().Name()); err != nil {
		return "", err
	}

	b, err := io.ReadAll(uc)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func validateExt(name string) error {
	ext := path.Ext(name)

	if ext == ".txt" || ext == ".doc" || ext == ".docs" || ext == ".ods" {
		return nil
	}
	return errors.New("unsuported file format")
}
