package main

import (
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

func (ui *UI) create_content() *fyne.Container {
	timer := ui.timer.get_widget()
	media_buttons := ui.create_media_buttons()
	time_buttons := ui.create_time_buttons()

	return container.New(layout.NewVBoxLayout(),
		timer,
		media_buttons,
		time_buttons,
	)
}

func (ui *UI) create_media_buttons() *fyne.Container {
	pr, _ := fyne.LoadResourceFromPath("resources/play-pause.png")
	s, _ := fyne.LoadResourceFromPath("resources/stop.png")

	return container.New(layout.NewGridLayout(2),
		widget.NewButtonWithIcon("", pr, func() { ui.timer.toggle() }),
		widget.NewButtonWithIcon("", s, func() { ui.timer.stop() }),
	)
}

func (ui *UI) create_time_buttons() *fyne.Container {
	return container.New(layout.NewGridLayout(3),
		widget.NewButton("25:00", func() { ui.timer.set(25 * time.Minute) }),
		widget.NewButton("15:00", func() { ui.timer.set(15 * time.Minute) }),
		widget.NewButton("10:00", func() { ui.timer.set(10 * time.Minute) }),
	)
}
