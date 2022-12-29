package main

import (
	"image/color"
	"time"
)

const (
	// Window
	WindowTitle  = "Pomodoro"
	WindowWidth  = 300
	WindowHeight = 150

	// Timer
	TimerDefaultTime = (25 * time.Minute)
)

var (
	// Theme
	TimerTextSize  = float32(100)
	TimerTextColor = color.White

	PauseTextColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}

	BackgroundColor = color.RGBA{R: 186, G: 73, B: 73, A: 150}

	DefaultBreaks = [...]time.Duration{
		5 * time.Minute,
		5 * time.Minute,
		15 * time.Minute,
	}
)
