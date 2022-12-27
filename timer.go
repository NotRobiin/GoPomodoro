package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Timer struct {
	ticker    *time.Ticker
	on_tick   func(*Timer)
	on_finish func(*Timer)
	paused    bool
	finished  bool
	tl        int
	text      *canvas.Text
}

func create_timer(on_tick func(*Timer), on_finish func(*Timer)) *Timer {
	t := new(Timer)

	t.ticker = time.NewTicker(1 * time.Second)
	t.on_tick = on_tick
	t.on_finish = on_finish
	t.paused = false
	t.finished = false
	t.text = canvas.NewText("", color.White)
	t.text.TextSize = TIMER_TEXT_SIZE

	return t
}

func (t *Timer) start() {
	go func() {
		for {
			select {
			case <-t.ticker.C:
				if t.paused {
					continue
				}

				t.tl--

				if t.tl < 0 {
					t.finished = true
					t.on_finish(t)
					t.stop()
					return
				}

				t.on_tick(t)
			}
		}
	}()
}

func (t *Timer) stop() {
	t.ticker.Stop()
}

func (t *Timer) pause() {
	t.paused = true
}

func (t *Timer) resume() {
	t.paused = false
}

func (t *Timer) toggle() {
	t.paused = !t.paused
}

func (t *Timer) set(tm time.Duration) {
	t.tl = int(tm.Seconds())
}

func (t *Timer) get_widget() *fyne.Container {
	return container.New(layout.NewCenterLayout(),
		t.text,
	)
}

func (t *Timer) show(ui *UI) {
	minutes := int(t.tl/60) % 60
	seconds := int(t.tl % 60)

	t.text.Text = fmt.Sprintf("%02d:%02d", minutes, seconds)
	t.text.Refresh()
}
