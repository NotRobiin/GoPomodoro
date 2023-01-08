package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

//go:embed notification.mp3
var notificationBytes []byte

var (
	ui       *UI
	isBreak  bool
	breakNum int
	sound    *Sound
)

func formatTime(tm time.Duration) string {
	s := int(tm.Seconds())
	minutes := int(s/60) % 60
	seconds := int(s % 60)

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func onMainTimerFinish() {
	isBreak = !isBreak
	newTime := parseTimeFromString(ui.p.StringWithFallback("timer", formatTime(DefaultSettings.timer)))
	s := BackgroundColor
	e := BackgroundColorBreak

	// Update breaks guts.
	if isBreak {
		newTime = DefaultBreaks[breakNum]
		ui.breaks[breakNum].enable()
		breakNum = (breakNum + 1) % len(DefaultBreaks)
	} else {
		s = BackgroundColorBreak
		e = BackgroundColor
	}

	// Update main timer.
	ui.timer.timer.stop()
	ui.timer.set(newTime)
	ui.timer.update()

	ui.bg.animate(s, e, BackgroundAnimationTime)

	if ui.p.BoolWithFallback("auto-start", DefaultSettings.autoStartEnabled) {
		ui.timer.started = true
		ui.timer.update()
		ui.timer.timer.countDown()
	} else {
		ui.timer.started = false
	}

	// Dim all the break widgets when we cycle back to the beginning.
	if !isBreak && breakNum == 0 {
		ui.disableBreaks()
	}

	if ui.p.BoolWithFallback("sound", DefaultSettings.soundEnabled) {
		sound.play(sound.decodeFromBytes(notificationBytes))
	}
}

func parseTimeFromString(s string) time.Duration {
	res := strings.Split(s, ":")
	min, _ := strconv.Atoi(res[0])
	sec, _ := strconv.Atoi(res[1])

	return time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
}

func main() {
	sound = new(Sound)
	sound.init()

	ui = new(UI)
	ui.app = app.NewWithID("gopomodoro.preferences")
	ui.p = ui.app.Preferences()

	ui.app.Settings().SetTheme(&newTheme{})
	ui.window = ui.app.NewWindow(WindowTitle)

	ui.createTray()

	ui.window.SetContent(ui.createContent())
	ui.window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	// ui.window.SetCloseIntercept(func() { ui.window.Hide() })
	ui.window.ShowAndRun()
}
