package main

import (
	"time"

	"fyne.io/fyne/v2/canvas"
)

type TimeWidget struct {
	widget *canvas.Text
	timer  *Timer
}

func createTimeWidget(finish func()) *TimeWidget {
	tw := new(TimeWidget)

	tw.widget = canvas.NewText("", TimerTextColor)
	tw.widget.TextSize = TimerTextSize
	tw.timer = createTimer(tw.onTick, finish)

	tw.set(TimerDefaultTime)
	tw.timer.countDown()

	return tw
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
	tw.update()
}

func (tw *TimeWidget) update() {
	tw.widget.Text = formatTime(tw.timer.tl)
	tw.widget.Refresh()
}
