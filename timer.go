package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Timer struct {
	ticker   *time.Ticker
	onTick   func(*Timer)
	onFinish func(*Timer)
	paused   bool
	finished bool
	tl       time.Duration
	text     *canvas.Text

	pauseWidget *canvas.Text
	pauseTicker *time.Ticker
	pauseTime   time.Duration
}

func create_timer(onTick func(*Timer), onFinish func(*Timer)) *Timer {
	t := new(Timer)

	t.ticker = nil
	t.onTick = onTick
	t.onFinish = onFinish
	t.paused = false
	t.finished = false

	t.text = canvas.NewText("", TimerTextColor)
	t.text.TextSize = TimerTextSize

	t.pauseWidget = canvas.NewText("", PauseTextColor)
	t.pauseTicker = nil
	t.pauseTime = 0 * time.Second

	return t
}

func (t *Timer) createTicker() {
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
					t.finished = true
					t.onFinish(t)
					t.stop()
					return
				}

				t.onTick(t)
			}
		}
	}()
}

func (t *Timer) stop() {
	t.ticker.Stop()
}

func (t *Timer) pause(visible bool) {
	t.paused = true

	if visible {
		t.pauseTicker = time.NewTicker(1 * time.Second)
		t.updatePauseTimer()

		go func() {
			for {
				select {
				case <-t.pauseTicker.C:
					t.pauseTime += time.Second
					t.updatePauseTimer()
				}
			}
		}()
	}
}

func (t *Timer) resume() {
	t.paused = false

	if t.pauseTicker != nil {
		t.pauseTicker.Stop()
		t.pauseTime = 0 * time.Second
		t.updatePauseTimer()
	}

	if t.ticker == nil {
		t.createTicker()
	}
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

	if t.ticker != nil {
		t.ticker.Stop()
	}

	t.createTicker()
}

func (t *Timer) updatePauseTimer() {
	if t.paused {
		t.pauseWidget.Text = fmt.Sprintf("Paused (%v)", formatTime(t.pauseTime))
	} else {
		t.pauseWidget.Text = ""
	}

	t.pauseWidget.Refresh()
}

func (t *Timer) show(ui *UI) {
	t.text.Text = formatTime(t.tl)
	t.text.Refresh()
}

func (t *Timer) getWidget() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewMaxLayout(),
				widget.NewButton("", func() {
					t.toggle()
				}),
				t.text,
			),
		),
		container.New(layout.NewCenterLayout(),
			t.pauseWidget,
		),
	)
}
