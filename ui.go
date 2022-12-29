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
	pauseText := canvas.NewText("Paused", PauseTextColor)

	return container.New(layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewMaxLayout(),
				widget.NewButton("", func() {
					ui.timer.toggle()

					if ui.timer.paused {
						pauseText.Text = "Paused"
					} else {
						pauseText.Text = ""
					}
				}),
				ui.timer.text,
			),
		),
		container.New(layout.NewCenterLayout(),
			pauseText,
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
