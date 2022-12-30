package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app    fyne.App
	window fyne.Window

	background *canvas.Rectangle
}

func (ui *UI) createContent() *fyne.Container {
	ui.background = canvas.NewRectangle(BackgroundColor)
	timers := ui.createTimerSegment()
	breaks := ui.createBreaksUI()

	return container.New(layout.NewMaxLayout(),
		ui.background,

		container.New(layout.NewVBoxLayout(),
			timers,
			breaks,
		),
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
	breaks = make([]*BreakWidget, len(DefaultBreaks))
	var widgets = make([]fyne.CanvasObject, len(DefaultBreaks))

	for i, v := range DefaultBreaks {
		breaks[i] = createBreakWidget(v)

		breaks[i].text.Text = formatTime(v)
		breaks[i].rect.SetMinSize(BreakRectSize)
		breaks[i].disable()

		widgets[i] = breaks[i].getWidget()
	}

	return container.New(layout.NewGridLayout(len(widgets)),
		widgets...,
	)
}
