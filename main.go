//go:generate fyne bundle -o data.go Icon.png
package main

import (
	"embed"

	"github.com/MaksimSvinin/gui-fresheye/internal/ui"
)

//go:embed README.md
var f embed.FS

func main() {
	ui.NewUI(f).Run()
}
