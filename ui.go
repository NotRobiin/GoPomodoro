package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app    fyne.App
	window fyne.Window
}

func (ui *UI) create_content(timer *Timer) *fyne.Container {
	timer_label := widget.NewLabel("")
	timer_text = binding.NewString()
	timer_label.Bind(timer_text)
	timer.show(timer_text)

	timer_container := container.New(layout.NewCenterLayout(),
		timer_label,
	)
	separator := widget.NewSeparator()

	pr, _ := fyne.LoadResourceFromPath("play-pause.png")
	s, _ := fyne.LoadResourceFromPath("stop.png")

	buttons := container.New(layout.NewGridLayout(2),
		widget.NewButtonWithIcon("", pr, func() { timer.toggle() }),
		widget.NewButtonWithIcon("", s, func() { timer.stop() }),
	)

	content := container.New(layout.NewVBoxLayout(), timer_container, separator, buttons)
	return content
}
