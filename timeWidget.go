package main

import (
	"time"

	"fyne.io/fyne/v2/canvas"
)

type TimeWidget struct {
	widget *canvas.Text
	timer  *Timer

	isBreak  bool
	breakNum int
}

func createTimeWidget() *TimeWidget {
	tw := new(TimeWidget)

	tw.widget = canvas.NewText("", TimerTextColor)
	tw.widget.TextSize = TimerTextSize
	tw.timer = createTimer(tw.onTick, tw.onFinish)
	tw.isBreak = false
	tw.breakNum = 0

	tw.set(TimerDefaultTime)
	tw.timer.start()

	return tw
}

func (tw *TimeWidget) onTick() {
	tw.update()
}

func (tw *TimeWidget) onFinish() {
	tw.isBreak = !tw.isBreak
	newTime := TimerDefaultTime

	if tw.isBreak {
		newTime = DefaultBreaks[tw.breakNum]
		breaks[tw.breakNum].enable()
		tw.breakNum = (tw.breakNum + 1) % len(DefaultBreaks)
	} else if tw.breakNum == 0 {
		for i := range breaks {
			breaks[i].disable()
		}
	}

	tw.timer.stop()
	tw.set(newTime)
	tw.timer.start()
}

func (tw *TimeWidget) toggle() {
	tw.timer.toggle()
}

func (tw *TimeWidget) set(tm time.Duration) {
	tw.timer.set(tm)
	tw.update()
}

func (tw *TimeWidget) update() {
	tw.widget.Text = formatTime(tw.timer.tl)
	tw.widget.Refresh()
}
