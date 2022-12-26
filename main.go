package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	timer_text binding.String
)

func on_tick(t *Timer) {
	t.show(timer_text)
}

func on_finish(t *Timer) {
}

func main() {
	timer := create_timer(on_tick, on_finish)
	timer.set(TEST_TIMER)
	timer.start()

	timer_label := widget.NewLabel("")
	timer_text = binding.NewString()
	timer_label.Bind(timer_text)
	timer.show(timer_text)

	myApp := app.New()
	myWindow := myApp.NewWindow(WINDOW_TITLE)
	myWindow.SetContent(timer_label)
	myWindow.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
