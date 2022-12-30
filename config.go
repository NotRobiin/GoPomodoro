package main

import (
	"image/color"
	"time"
)

const (
	// Window
	WindowTitle  = "Pomodoro"
	WindowWidth  = 300
	WindowHeight = 200

	// Timer
	TimerDefaultTime = (3 * time.Second)
)

var (
	// Theme
	TimerTextSize  = float32(100)
	TimerTextColor = color.White

	PauseTextColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}

	BackgroundColor      = color.RGBA{R: 186, G: 73, B: 73, A: 150}
	BackgroundColorBreak = color.RGBA{R: 102, G: 153, B: 255, A: 150}

	DefaultBreaks = [...]time.Duration{
		5 * time.Second,
		10 * time.Second,
		15 * time.Second,
	}
)
