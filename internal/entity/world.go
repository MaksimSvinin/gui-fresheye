package entity

import "image/color"

type WorldInfo struct {
	Count   int
	Indexes []WorldIndex
	World   string
	Color   color.Color
}

type WorldIndex struct {
	From int
	To   int
}

type WorldEndInfo struct {
	To       int
	World    string
	Color    color.Color
	CloseLoc bool
}
