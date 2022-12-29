package main

import (
	"fmt"

	"fyne.io/fyne/v2/canvas"
)

type PauseWidget struct {
	widget *canvas.Text
	timer  *Timer
}

func createPauseWidget() *PauseWidget {
	pw := new(PauseWidget)
	pw.widget = canvas.NewText("", PauseTextColor)
	pw.timer = createTimer(func() { pw.update() }, func() {})

	return pw
}

func (pw *PauseWidget) pause() {
	pw.timer.pause()
	pw.update()
}

func (pw *PauseWidget) unpause() {
	pw.timer.unpause()
	pw.update()
}

func (pw *PauseWidget) toggle() {
	pw.timer.toggle()
	pw.update()
}

func (pw *PauseWidget) update() {
	if pw.timer.paused {
		pw.widget.Text = fmt.Sprintf("Paused (%v)", formatTime(pw.timer.tl))
	} else {
		pw.widget.Text = ""
	}

	pw.widget.Refresh()
}
