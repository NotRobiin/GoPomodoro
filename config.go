package main

import (
	"time"
)

const (
	// Window
	WindowTitle  = "Pomodoro"
	WindowWidth  = 300
	WindowHeight = 300

	// Timer
	TimerDefaultTime = (25 * time.Minute)
	TimerTextSize    = 75.0
)

var (
	TimerDefaultTimes = [...]time.Duration{
		25 * time.Minute,
		15 * time.Minute,
		10 * time.Minute,
		5 * time.Minute,
	}

	DefaultBreaks = [...]time.Duration{
		5 * time.Minute,
		5 * time.Minute,
		15 * time.Minute,
	}
)
