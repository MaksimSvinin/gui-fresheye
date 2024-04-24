package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomText struct {
	style widget.RichTextStyle
	text  string

	r, g, b uint8
	from    int
	to      int
}

func NewCustomText(text string, r, g, b uint8, from, to int) *CustomText {
	return &CustomText{
		r: r,
		g: g,
		b: b,

		text:  text,
		style: widget.RichTextStyleCodeBlock,

		from: from,
		to:   to,
	}
}

func (t *CustomText) color() color.Color {
	return color.RGBA{t.r, t.g, t.b, 255}
}

func (t *CustomText) From() int {
	return t.from
}

func (t *CustomText) To() int {
	return t.to
}

// Inline should return true if this text can be included within other elements, or false if it creates a new block.
func (t *CustomText) Inline() bool {
	return t.style.Inline
}

// Textual returns the content of this segment rendered to plain text.
func (t *CustomText) Textual() string {
	return t.text
}

// Visual returns the graphical elements required to render this segment.
func (t *CustomText) Visual() fyne.CanvasObject {
	obj := canvas.NewText(t.text, t.color())

	t.Update(obj)
	return obj
}

// Update applies the current state of this text segment to an existing visual.
func (t *CustomText) Update(o fyne.CanvasObject) {
	//nolint:errcheck // copy
	obj := o.(*canvas.Text)
	obj.Text = t.text
	obj.Color = t.color()
	obj.Alignment = t.style.Alignment
	obj.TextStyle = t.style.TextStyle
	obj.TextSize = t.size()
	obj.Refresh()
}

// Select tells the segment that the user is selecting the content between the two positions.
func (t *CustomText) Select(_, _ fyne.Position) {
	// no-op: this will be added when we progress to editor
}

// SelectedText should return the text representation of any content currently selected through the Select call.
func (t *CustomText) SelectedText() string {
	// no-op: this will be added when we progress to editor
	return ""
}

// Unselect tells the segment that the user is has cancelled the previous selection.
func (t *CustomText) Unselect() {
	// no-op: this will be added when we progress to editor
}

func (t *CustomText) size() float32 {
	if t.style.SizeName != "" {
		return fyne.CurrentApp().Settings().Theme().Size(t.style.SizeName)
	}

	return theme.TextSize()
}
