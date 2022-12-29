package main

import (
	"time"

	"fyne.io/fyne/v2/canvas"
)

type Timer struct {
	ticker   *time.Ticker
	onTick   func(*Timer)
	onFinish func(*Timer)
	paused   bool
	tl       time.Duration
	text     *canvas.Text
}

func create_timer(onTick func(*Timer), onFinish func(*Timer)) *Timer {
	t := new(Timer)

	t.ticker = nil
	t.onTick = onTick
	t.onFinish = onFinish
	t.paused = false

	t.text = canvas.NewText("", TimerTextColor)
	t.text.TextSize = TimerTextSize

	return t
}

func (t *Timer) startTicker() {
	t.ticker = time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-t.ticker.C:
				if t.paused {
					continue
				}

				t.tl -= time.Second

				if t.tl < 0 {
					t.ticker.Stop()
					t.onFinish(t)
					return
				}

				t.onTick(t)
			}
		}
	}()
}

func (t *Timer) pause(visible bool) {
	t.paused = true
}

func (t *Timer) resume() {
	t.paused = false
}

func (t *Timer) toggle() {
	t.paused = !t.paused

	if t.paused {
		t.pause(true)
	} else {
		t.resume()
	}
}

func (t *Timer) set(tm time.Duration) {
	t.tl = tm

	t.text.Text = formatTime(t.tl)
	t.text.Refresh()
}

func (t *Timer) show(ui *UI) {
	t.text.Text = formatTime(t.tl)
	t.text.Refresh()
}
