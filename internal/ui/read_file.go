package ui

import (
	"errors"
	"io"
	"path"

	"fyne.io/fyne/v2"
	"github.com/gelsrc/go-charset"
)

const (
	extTxt = ".txt"
)

func readFile(uc fyne.URIReadCloser, win1251 bool) (string, error) {
	if uc == nil {
		return "", errors.New("not select file")
	}
	defer uc.Close()
	if err := validateExt(uc.URI().Name()); err != nil {
		return "", err
	}

	return readTxt(uc, win1251)
}

func readTxt(uc fyne.URIReadCloser, win1251 bool) (string, error) {
	b, err := io.ReadAll(uc)
	if err != nil {
		return "", err
	}

	if win1251 {
		return string(charset.Cp1251BytesToRunes(b)), nil
	}

	return string(b), nil
}

func validateExt(name string) error {
	ext := path.Ext(name)

	if ext == extTxt {
		return nil
	}
	return errors.New("unsuported file format")
}
