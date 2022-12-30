package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/canvas"
)

type Background struct {
	widget *canvas.Rectangle
}

func createBackground(c color.Color) *Background {
	return &Background{
		widget: canvas.NewRectangle(c),
	}
}

func (bg *Background) animate(s, e color.Color, t time.Duration) {
	canvas.NewColorRGBAAnimation(s, e, t, func(c color.Color) {
		bg.widget.FillColor = c
		bg.widget.Refresh()
	}).Start()
}
