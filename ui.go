package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
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
	timer := ui.timer.getWidget()
	mediaButtons := ui.createMediaButtons()
	timeButtons := ui.createTimeButtons()

	return container.New(layout.NewVBoxLayout(),
		timer,
		mediaButtons,
		timeButtons,
	)
}

func (ui *UI) createMediaButtons() *fyne.Container {
	pr, _ := fyne.LoadResourceFromPath("resources/play-pause.png")
	s, _ := fyne.LoadResourceFromPath("resources/stop.png")

	return container.New(layout.NewGridLayout(2),
		widget.NewButtonWithIcon("", pr, func() { ui.timer.toggle() }),
		widget.NewButtonWithIcon("", s, func() { ui.timer.stop() }),
	)
}

func (ui *UI) createTimeButtons() *fyne.Container {
	var buttons = make([]fyne.CanvasObject, len(TimerDefaultTimes))

	for i, v := range TimerDefaultTimes {
		s := int(v.Seconds())
		minutes := int(s/60) % 60
		seconds := int(s % 60)
		f := func(v time.Duration) func() {
			return func() { ui.timer.set(v) }
		}

		buttons[i] = widget.NewButton(fmt.Sprintf("%02d:%02d", minutes, seconds), f(v))
	}

	return container.New(layout.NewGridLayout(len(TimerDefaultTimes)),
		buttons...,
	)
}
