package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	ui *UI
)

func on_tick(timer *Timer) {
	timer.show(ui)
}

func on_finish(timer *Timer) {
}

func main() {
	ui = new(UI)
	ui.app = app.New()

	timer := create_timer(on_tick, on_finish)
	timer.set(TIMER_DEFAULT_TIME)
	timer.start()

	ui.timer = timer
	ui.window = ui.app.NewWindow(WINDOW_TITLE)
	ui.window.SetContent(ui.create_content())
	ui.window.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	ui.window.ShowAndRun()
}
