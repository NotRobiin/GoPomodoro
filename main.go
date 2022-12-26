package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	timer_text binding.String
)

func on_tick(timer *Timer) {
	timer.show(timer_text)
}

func on_finish(timer *Timer) {
}

func create_timer_label() *widget.Label {
	timer_label := widget.NewLabel("")
	timer_text = binding.NewString()
	timer_label.Bind(timer_text)

	return timer_label
}

func create_content(timer *Timer) *fyne.Container {
	timer_label := create_timer_label()
	timer.show(timer_text)

	separator := widget.NewSeparator()

	buttons := container.New(layout.NewGridLayout(3),
		widget.NewButton("Pause", func() { timer.pause() }),
		widget.NewButton("Resume", func() { timer.resume() }),
		widget.NewButton("Stop", func() { timer.stop() }),
	)

	content := container.New(layout.NewVBoxLayout(), timer_label, separator, buttons)
	return content
}

func main() {
	timer := create_timer(on_tick, on_finish)
	timer.set(TEST_TIMER)
	timer.start()

	gui := app.New()
	window := gui.NewWindow(WINDOW_TITLE)
	window.SetContent(create_content(timer))
	window.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	window.SetMaster()
	window.CenterOnScreen()
	window.ShowAndRun()
}
