package main

import (
	"io"
	"os"
	"sort"

	"fyne.io/fyne/v2/widget"

	"gitlab.com/opennota/fresheye"

	"github.com/MaksimSvinin/gui-fresheye/ui"
)

func main() {
	excludeProperNames := false
	sensitivityThreshold := 10
	contextSize := 10

	f, _ := os.Open("/home/max/go/src/github.com/MaksimSvinin/gui-fresheye/t.txt")
	defer f.Close()

	tb, _ := io.ReadAll(f)
	text := string(tb)

	it := fresheye.NewSimpleWordIterator(text)
	bad := fresheye.Check(it,
		fresheye.ContextSize(contextSize),
		fresheye.SensitivityThreshold(sensitivityThreshold),
		fresheye.ExcludeProperNames(excludeProperNames),
	)
	colorizer := fresheye.NewColorizer(sensitivityThreshold)
	colorizer.Colorize(bad)
	descriptors := fresheye.WordDescriptors(bad)

	sort.Slice(descriptors, func(i, j int) bool {
		return descriptors[i].Start() < descriptors[j].Start()
	})

	t := make([]widget.RichTextSegment, 0, len(descriptors))

	// i := 0
	for _, d := range descriptors {
		c := colorizer.Color(d)
		from, to := d.Start(), d.End()
		world := text[from:to]

		t = append(t, ui.NewCustomText(world, c.R, c.G, c.B, from, to))
		// i = to
	}
	// fmt.Println(text[i:])

	app := ui.NewUI(t[0:100])

	app.Run()
}
