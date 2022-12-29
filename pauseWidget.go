package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/canvas"
)

type PauseWidget struct {
	widget *canvas.Text
	ticker *time.Ticker
	time   time.Duration
	paused bool
}

func createPauseWidget() *PauseWidget {
	return &PauseWidget{
		widget: canvas.NewText("", PauseTextColor),
		ticker: nil,
		time:   0 * time.Second,
	}
}

func (pw *PauseWidget) startTimer() {
	pw.ticker = time.NewTicker(1 * time.Second)
	pw.paused = true
	pw.updatePauseTimer()

	go func() {
		for {
			select {
			case <-pw.ticker.C:
				pw.time += time.Second
				pw.updatePauseTimer()
			}
		}
	}()
}

func (pw *PauseWidget) stopTimer() {
	if pw.ticker != nil {
		pw.ticker.Stop()
	}

	pw.time = 0 * time.Second
	pw.paused = false
	pw.updatePauseTimer()
}

func (pw *PauseWidget) updatePauseTimer() {
	if pw.paused {
		pw.widget.Text = fmt.Sprintf("Paused (%v)", formatTime(pw.time))
	} else {
		pw.widget.Text = ""
	}

	pw.widget.Refresh()
}
