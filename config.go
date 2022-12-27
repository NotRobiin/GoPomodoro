package main

import "time"

const (
	// Window
	WINDOW_TITLE  = "Pomodoro"
	WINDOW_WIDTH  = 300
	WINDOW_HEIGHT = 300

	// Timer
	TIMER_DEFAULT_TIME = (25 * time.Minute)
	TIMER_TEXT_SIZE    = 75.0
)

var (
	TIMER_DEFAULT_TIMES = [...]time.Duration{
		25 * time.Minute,
		15 * time.Minute,
		10 * time.Minute,
		5 * time.Minute,
	}
)
