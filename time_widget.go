package main

import (
	"time"

	"fyne.io/fyne/v2/canvas"
)

type TimeWidget struct {
	widget  *canvas.Text
	timer   *Timer
	started bool
}

func createTimeWidget(onFinish func()) *TimeWidget {
	tw := new(TimeWidget)

	tw.widget = canvas.NewText("", TimerTextColor)
	tw.widget.Text = TimerStartText
	tw.widget.TextSize = TimerTextSize
	tw.timer = createTimer(tw.onTick, onFinish)
	tw.timer.set(settings.timer)

	return tw
}

func (tw *TimeWidget) start() {
	tw.started = true
	tw.update()
	tw.timer.countDown()
}

func (tw *TimeWidget) onTick() {
	tw.update()
}

func (tw *TimeWidget) toggle() {
	tw.timer.toggle()

	if tw.timer.paused {
		tw.widget.Color = TimerTextColorPaused
	} else {
		tw.widget.Color = TimerTextColor
	}

	tw.widget.Refresh()
}

func (tw *TimeWidget) set(tm time.Duration) {
	tw.timer.set(tm)

	if tw.started {
		tw.update()
	}
}

func (tw *TimeWidget) update() {
	tw.widget.Text = formatTime(tw.timer.tl)
	tw.widget.Refresh()
}
