package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type BreakWidget struct {
	rect *canvas.Rectangle
	text *canvas.Text
	tm   time.Duration
}

func createBreakWidget(t time.Duration) *BreakWidget {
	return &BreakWidget{
		rect: canvas.NewRectangle(color.White),
		text: canvas.NewText("", color.White),
		tm:   t,
	}
}

func (bw *BreakWidget) getWidget() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		bw.rect,
		bw.text,
	)
}

func (bw *BreakWidget) disable() {
	bw.text.Color = color.RGBA{R: 80, G: 80, B: 80, A: 75}
	bw.rect.FillColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}
	bw.rect.Refresh()
	bw.text.Refresh()
}

func (bw *BreakWidget) enable() {
	bw.text.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	bw.rect.FillColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	bw.rect.Refresh()
	bw.text.Refresh()
}
