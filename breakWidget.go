package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/canvas"
)

type BreakWidget struct {
	rect *canvas.Rectangle
	text *canvas.Text
	tm   time.Duration
}

func createBreakWidget(t time.Duration, c color.Color) *BreakWidget {
	return &BreakWidget{
		rect: canvas.NewRectangle(c),
		text: canvas.NewText("", color.White),
		tm:   t,
	}
}
