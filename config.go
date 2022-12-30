package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
)

const (
	// Window
	WindowTitle  = "Pomodoro"
	WindowWidth  = 300
	WindowHeight = 200

	// Timer
	TimerDefaultTime = (25 * time.Minute)
)

var (
	// Theme
	TimerTextSize        = float32(100)
	TimerTextColor       = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	TimerTextColorPaused = color.RGBA{R: 80, G: 80, B: 80, A: 75}

	PauseTextColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	BackgroundColor      = color.RGBA{R: 186, G: 73, B: 73, A: 150}
	BackgroundColorBreak = color.RGBA{R: 102, G: 153, B: 255, A: 150}

	BreakDisabledTextColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}
	BreakDisabledRectColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}

	BreakEnabledTextColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	BreakEnabledRectColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	BreakRectSize = fyne.NewSize(3, 3)

	BackgroundAnimationTime = 450 * time.Millisecond

	DefaultBreaks = [...]time.Duration{
		5 * time.Minute,
		5 * time.Minute,
		15 * time.Minute,
	}
)
