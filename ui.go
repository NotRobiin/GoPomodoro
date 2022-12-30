package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app    fyne.App
	window fyne.Window
}

func (ui *UI) createContent() *fyne.Container {
	timers := ui.createTimerSegment()
	breaks := ui.createBreaksUI()

	return container.New(layout.NewVBoxLayout(),
		timers,
		breaks,
	)
}

func (ui *UI) createTimerSegment() *fyne.Container {
	timeWidget := createTimeWidget()
	pauseWidget := createPauseWidget()
	toggleButtonWidget := widget.NewButton("", func() {
		timeWidget.toggle()
		pauseWidget.toggle()
	})

	return container.New(layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewMaxLayout(),
				toggleButtonWidget,
				timeWidget.widget,
			),
		),
		container.New(layout.NewCenterLayout(),
			pauseWidget.widget,
		),
	)
}

func (ui *UI) createBreaksUI() *fyne.Container {
	var breaks = make([]fyne.CanvasObject, len(DefaultBreaks))

	for i, v := range DefaultBreaks {
		b := createBreakWidget(v, color.White)

		b.text.Text = formatTime(v)

		breaks[i] = container.New(layout.NewVBoxLayout(),
			b.rect,
			b.text,
		)
	}

	return container.New(layout.NewGridLayout(len(breaks)),
		breaks...,
	)
}
