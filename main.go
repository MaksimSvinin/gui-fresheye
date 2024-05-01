//go:generate fyne bundle -o data.go Icon.png
package main

import (
	"github.com/MaksimSvinin/gui-fresheye/internal/ui"
)

func main() {
	ui.NewUI().Run()
}
