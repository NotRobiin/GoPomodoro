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
	TimerDefaultTime = (5 * time.Second)
)

var (
	// Theme
	TimerTextSize        = float32(100)
	TimerStartText       = "Start!"
	TimerTextColor       = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	TimerTextColorPaused = color.RGBA{R: 80, G: 80, B: 80, A: 75}

	// Pause
	PauseTextColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	// Background
	BackgroundColor         = color.RGBA{R: 186, G: 73, B: 73, A: 150}
	BackgroundColorBreak    = color.RGBA{R: 102, G: 153, B: 255, A: 150}
	BackgroundAnimationTime = 450 * time.Millisecond
	OverlayBackgroundColor  = color.RGBA{R: 186, G: 73, B: 73, A: 237}

	// Breaks
	BreakDisabledTextColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}
	BreakDisabledRectColor = color.RGBA{R: 80, G: 80, B: 80, A: 75}

	BreakEnabledTextColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	BreakEnabledRectColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}

	BreakRectSize = fyne.NewSize(3, 3)

	// Notifications
	NotificationButtonMultiplier = 1.5 // Default size (theme.IconInlineSize()) * value
	NotificationSound            = "./resources/notification.mp3"

	DefaultBreaks = [...]time.Duration{
		5 * time.Second,
		5 * time.Second,
		15 * time.Second,
	}
)
