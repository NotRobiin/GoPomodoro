package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	app      fyne.App
	window   fyne.Window
	tray     desktop.App
	settings *SettingsWidget

	bg     *Background
	timer  *TimeWidget
	breaks []*BreakWidget
}

func (ui *UI) createTray() {
	if desk, ok := ui.app.(desktop.App); ok {
		ui.tray = desk

		m := fyne.NewMenu(WindowTitle,
			fyne.NewMenuItem("Show", func() {
				ui.window.Show()
			}))

		ui.tray.SetSystemTrayMenu(m)
	}
}

func (ui *UI) createContent() *fyne.Container {
	ui.bg = createBackground(BackgroundColor)
	timers := ui.createTimerSegment()
	breaks := ui.createBreaksUI()
	ui.settings = ui.createSettings()

	bStop := widget.NewButtonWithIcon("", theme.MediaStopIcon(), func() {
		isBreak = false
		breakNum = 0
		t := parseTimeFromString(pref.StringWithFallback("timer", formatTime(TimerDefaultTime)))

		ui.timer.restart(t)
		ui.disableBreaks()
		ui.bg.animate(ui.bg.widget.FillColor, BackgroundColor, BackgroundAnimationTime)
	})
	bSkip := widget.NewButtonWithIcon("", theme.MediaSkipNextIcon(), func() {
		ui.timer.skip()
	})

	return container.New(layout.NewMaxLayout(),
		ui.bg.widget,
		ui.settings.overlay,
		container.New(layout.NewVBoxLayout(),
			container.New(layout.NewHBoxLayout(), layout.NewSpacer(), bStop, bSkip, ui.settings.toggleButton),
		),

		container.New(layout.NewVBoxLayout(),
			layout.NewSpacer(),
			timers,
			breaks,
		),
	)
}

func (ui *UI) createSettings() *SettingsWidget {
	s := &SettingsWidget{title: "Settings"}
	s.create(ui.window.Canvas(), func() { s.toggle() })

	s.toggleButton = widget.NewButtonWithIcon("", theme.SettingsIcon(), func() { s.toggle() })
	s.toggleButton.Move(fyne.NewPos(WindowWidth-theme.IconInlineSize(), 0))
	s.toggleButton.Resize(fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize()))

	// Sound
	sSound := widget.NewCheck("", func(v bool) { pref.SetBool("sound", v) })
	sSound.SetChecked(pref.BoolWithFallback("sound", DefaultSettings.soundEnabled))
	s.add("Sound", container.New(layout.NewMaxLayout(), sSound), layout.NewHBoxLayout(), true)

	// Auto start
	sAutoStart := widget.NewCheck("", func(v bool) { pref.SetBool("auto-start", v) })
	sAutoStart.SetChecked(pref.BoolWithFallback("auto-start", DefaultSettings.autoStartEnabled))
	s.add("Auto-start", container.New(layout.NewMaxLayout(), sAutoStart), layout.NewHBoxLayout(), true)

	// Timer
	var sTimer *widget.CheckGroup
	var times []string

	for _, v := range TimerOptions {
		times = append(times, formatTime(v))
	}

	sTimer = widget.NewCheckGroup(times, func(s []string) {
		if len(s) <= 1 {
			return
		}

		chosen := s[len(s)-1]
		newTime := parseTimeFromString(chosen)

		pref.SetString("timer", chosen)

		if !ui.timer.started {
			ui.timer.set(newTime)
		}

		sTimer.SetSelected([]string{chosen})
	})
	sTimer.Horizontal = true
	sTimer.SetSelected([]string{pref.StringWithFallback("timer", formatTime(DefaultSettings.timer))})
	s.add("Timer", container.New(layout.NewHBoxLayout(), sTimer), layout.NewHBoxLayout(), true)

	s.hide()

	return s
}

func (ui *UI) createTimerSegment() *fyne.Container {
	t := parseTimeFromString(pref.StringWithFallback("timer", formatTime(DefaultSettings.timer)))
	ui.timer = createTimeWidget(onMainTimerFinish, t)
	pauseWidget := createPauseWidget()
	toggleButtonWidget := widget.NewButton("", func() {
		if ui.timer.started {
			ui.timer.toggle()
			pauseWidget.toggle()
		} else {
			ui.timer.start()
		}
	})

	return container.New(layout.NewVBoxLayout(),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewMaxLayout(),
				toggleButtonWidget,
				ui.timer.widget,
			),
		),
		container.New(layout.NewCenterLayout(),
			pauseWidget.widget,
		),
	)
}

func (ui *UI) createBreaksUI() *fyne.Container {
	ui.breaks = make([]*BreakWidget, len(DefaultBreaks))
	var widgets = make([]fyne.CanvasObject, len(DefaultBreaks))

	for i, v := range DefaultBreaks {
		ui.breaks[i] = createBreakWidget(v)

		ui.breaks[i].text.Text = formatTime(v)
		ui.breaks[i].rect.SetMinSize(BreakRectSize)
		ui.breaks[i].disable()

		widgets[i] = ui.breaks[i].getWidget()
	}

	return container.New(layout.NewGridLayout(len(widgets)),
		widgets...,
	)
}

func (ui *UI) disableBreaks() {
	for i := range ui.breaks {
		ui.breaks[i].disable()
	}
}
