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
	timer  *Timer

	timer_text binding.String
}

func (ui *UI) create_content() *fyne.Container {
	timer := ui.create_timer_label()
	separator := widget.NewSeparator()
	buttons := ui.create_buttons()

	return container.New(layout.NewVBoxLayout(),
		timer,
		separator,
		buttons,
	)
}

func (ui *UI) create_buttons() *fyne.Container {
	pr, _ := fyne.LoadResourceFromPath("resources/play-pause.png")
	s, _ := fyne.LoadResourceFromPath("resources/stop.png")

	return container.New(layout.NewGridLayout(2),
		widget.NewButtonWithIcon("", pr, func() { ui.timer.toggle() }),
		widget.NewButtonWithIcon("", s, func() { ui.timer.stop() }),
	)
}

func (ui *UI) create_timer_label() *fyne.Container {
	timer_label := widget.NewLabel("")
	ui.timer_text = binding.NewString()
	timer_label.Bind(ui.timer_text)
	ui.timer.show(ui.timer_text)

	return container.New(layout.NewCenterLayout(),
		timer_label,
	)
}
