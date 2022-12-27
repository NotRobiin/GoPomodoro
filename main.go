package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	ui *UI
)

func onTick(timer *Timer) {
	timer.show(ui)
}

func onFinish(timer *Timer) {
}

func main() {
	ui = new(UI)
	ui.app = app.New()

	timer := create_timer(onTick, onFinish)
	timer.set(TimerDefaultTime)
	timer.start()

	ui.timer = timer
	ui.window = ui.app.NewWindow(WindowTitle)
	ui.window.SetContent(ui.createContent())
	ui.window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	ui.window.ShowAndRun()
}
