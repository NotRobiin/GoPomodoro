package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/data/binding"
)

type Timer struct {
	ticker    *time.Ticker
	on_tick   func(*Timer)
	on_finish func(*Timer)
	paused    bool
	finished  bool
	tl        int
}

func create_timer(on_tick func(*Timer), on_finish func(*Timer)) *Timer {
	t := new(Timer)

	t.ticker = time.NewTicker(1 * time.Second)
	t.on_tick = on_tick
	t.on_finish = on_finish
	t.paused = false
	t.finished = false

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

func (t *Timer) set(tm time.Duration) {
	t.tl = int(tm.Seconds())
}

func (t *Timer) show(bind binding.String) {
	minutes := int(t.tl/60) % 60
	seconds := int(t.tl % 60)

	bind.Set(fmt.Sprintf("%02d:%02d", minutes, seconds))
}
