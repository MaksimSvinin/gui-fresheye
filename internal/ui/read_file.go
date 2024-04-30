package ui

import (
	"errors"
	"io"
	"path"

	"fyne.io/fyne/v2"
	"github.com/lu4p/cat"
)

const (
	extTxt  = ".txt"
	extDocx = ".docx"
	extOdt  = ".odt"
	extRtf  = ".rtf"
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
	return cat.FromBytes(b)
}

func validateExt(name string) error {
	ext := path.Ext(name)

	if ext == extTxt || ext == extDocx || ext == extOdt || ext == extRtf {
		return nil
	}
	return errors.New("unsuported file format")
}
