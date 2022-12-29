package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app    fyne.App
	window fyne.Window
	timer  *Timer
}

func (ui *UI) createContent() *fyne.Container {
	timer := ui.createTimer()
	breaks := ui.createBreaksUI()

	return container.New(layout.NewGridLayoutWithRows(2),
		timer,
		breaks,
	)
}

func (ui *UI) createTimer() *fyne.Container {
	return container.New(layout.NewCenterLayout(),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewMaxLayout(),
				widget.NewButton("", func() {
					ui.timer.toggle()
				}),
				ui.timer.text,
			),
		),
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

	return container.New(layout.NewGridLayout(3),
		breaks...,
	)
}
