package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomSegment struct {
	Style widget.RichTextStyle
	Text  string
	Color color.Color
}

// Inline should return true if this text can be included within other elements, or false if it creates a new block.
func (t *CustomSegment) Inline() bool {
	return true
}

// Textual returns the content of this segment rendered to plain text.
func (t *CustomSegment) Textual() string {
	return t.Text
}

// Visual returns the graphical elements required to render this segment.
func (t *CustomSegment) Visual() fyne.CanvasObject {
	obj := canvas.NewText(t.Text, t.color())

	t.Update(obj)
	return obj
}

// Update applies the current state of this text segment to an existing visual.
func (t *CustomSegment) Update(o fyne.CanvasObject) {
	obj := o.(*canvas.Text)
	obj.Text = t.Text
	obj.Color = t.color()
	obj.Alignment = t.Style.Alignment
	obj.TextStyle = t.Style.TextStyle
	obj.TextSize = t.size()
	obj.Refresh()
}

// Select tells the segment that the user is selecting the content between the two positions.
func (t *CustomSegment) Select(_, _ fyne.Position) {
	// no-op: this will be added when we progress to editor
}

// SelectedText should return the text representation of any content currently selected through the Select call.
func (t *CustomSegment) SelectedText() string {
	// no-op: this will be added when we progress to editor
	return ""
}

// Unselect tells the segment that the user is has cancelled the previous selection.
func (t *CustomSegment) Unselect() {
	// no-op: this will be added when we progress to editor
}

func (t *CustomSegment) color() color.Color {
	return t.Color
}

func (t *CustomSegment) size() float32 {
	if t.Style.SizeName != "" {
		return fyne.CurrentApp().Settings().Theme().Size(t.Style.SizeName)
	}

	return theme.TextSize()
}
