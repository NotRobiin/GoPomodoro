package main

import (
	"time"
)

type Timer struct {
	ticker   *time.Ticker
	onTick   func()
	onFinish func()
	tl       time.Duration
	paused   bool
}

func createTimer(onTick func(), onFinish func()) *Timer {
	return &Timer{
		ticker:   nil,
		onTick:   onTick,
		onFinish: onFinish,
		tl:       0,
		paused:   false,
	}
}

func (t *Timer) pause() {
	t.paused = true
}

func (t *Timer) unpause() {
	t.paused = false
}

func (t *Timer) toggle() {
	if t.paused {
		t.unpause()
	} else {
		t.pause()
	}
}

func (t *Timer) set(tm time.Duration) {
	t.tl = tm
}

func (t *Timer) create() {
	t.ticker = time.NewTicker(1 * time.Second)
}

func (t *Timer) stop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}

	t.ticker = nil
}

func (t *Timer) countDown() {
	if t.ticker == nil {
		t.create()
	}

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
					t.onFinish()
					return
				}

				t.onTick()
			}
		}
	}()
}

func (t *Timer) countUp() {
	if t.ticker == nil {
		t.create()
	}

	go func() {
		for {
			select {
			case <-t.ticker.C:
				if t.paused {
					continue
				}

				t.tl += time.Second
				t.onTick()
			}
		}
	}()
}
