package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/canvas"
)

type PauseWidget struct {
	widget  *canvas.Text
	timer   *Timer
	enabled bool
}

func createPauseWidget() *PauseWidget {
	pw := new(PauseWidget)
	pw.widget = canvas.NewText("", PauseTextColor)
	pw.timer = createTimer(func() { pw.update() }, func() {})
	pw.enabled = false

	return pw
}

func (pw *PauseWidget) toggle() {
	pw.enabled = !pw.enabled

	if pw.enabled {
		pw.timer.unpause()
		pw.timer.set(0 * time.Second)
		pw.timer.countUp()
	} else {
		pw.timer.pause()
		pw.timer.stop()
	}

	pw.update()
}

func (pw *PauseWidget) update() {
	if pw.enabled {
		pw.widget.Text = fmt.Sprintf("PAUSED (%v)", formatTime(pw.timer.tl))
	} else {
		pw.widget.Text = ""
	}

	pw.widget.Refresh()
}
