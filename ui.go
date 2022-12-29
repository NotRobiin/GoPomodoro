package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type UI struct {
	app    fyne.App
	window fyne.Window
	timer  *Timer
}

func (ui *UI) createContent() *fyne.Container {
	timer := ui.timer.getWidget()
	breaks := ui.createBreaksUI()

	return container.New(layout.NewVBoxLayout(),
		timer,
		breaks,
	)
}

func (ui *UI) createBreaksUI() *fyne.Container {
	var breaks = make([]fyne.CanvasObject, len(DefaultBreaks))

	for i, v := range DefaultBreaks {
		breaks[i] = container.New(layout.NewVBoxLayout(),
			canvas.NewRectangle(color.White),
			canvas.NewText(formatTime(v), color.White),
		)
	}

	return container.New(layout.NewGridLayout(len(breaks)),
		breaks...,
	)
}
