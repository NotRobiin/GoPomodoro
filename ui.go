package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app    fyne.App
	window fyne.Window

	bg    *Background
	timer *TimeWidget
}

func (ui *UI) createContent() *fyne.Container {
	ui.bg = createBackground(BackgroundColor)
	timers := ui.createTimerSegment()
	breaks := ui.createBreaksUI()

	return container.New(layout.NewMaxLayout(),
		ui.bg.widget,

		container.New(layout.NewVBoxLayout(),
			timers,
			breaks,
		),
	)
}

func (ui *UI) createTimerSegment() *fyne.Container {
	ui.timer = createTimeWidget(onMainTimerFinish)
	pauseWidget := createPauseWidget()
	toggleButtonWidget := widget.NewButton("", func() {
		ui.timer.toggle()
		pauseWidget.toggle()
	})

	return container.New(layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewMaxLayout(),
				toggleButtonWidget,
				ui.timer.widget,
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
