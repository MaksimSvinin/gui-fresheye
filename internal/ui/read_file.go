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

	notFoundFile     = "не выбран фаил"
	notSupportFormat = "формат файла не поддерживается"
)

func readFile(uc fyne.URIReadCloser, win1251 bool) (string, error) {
	if uc == nil {
		return "", errors.New(notFoundFile)
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
	return errors.New(notSupportFormat)
}
