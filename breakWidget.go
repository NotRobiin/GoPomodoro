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
	bw.text.Color = BreakDisabledTextColor
	bw.rect.FillColor = BreakDisabledRectColor
	bw.rect.Refresh()
	bw.text.Refresh()
}

func (bw *BreakWidget) enable() {
	bw.text.Color = BreakEnabledTextColor
	bw.rect.FillColor = BreakEnabledRectColor
	bw.rect.Refresh()
	bw.text.Refresh()
}
