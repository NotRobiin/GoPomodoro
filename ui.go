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
	app    fyne.App
	window fyne.Window
	tray   desktop.App

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
	soundToggle := ui.createNotificationButton()

	return container.New(layout.NewMaxLayout(),
		ui.bg.widget,
		container.NewWithoutLayout(soundToggle),

		container.New(layout.NewVBoxLayout(),
			timers,
			breaks,
		),
	)
}

func (ui *UI) createTimerSegment() *fyne.Container {
	ui.timer = createTimeWidget(onMainTimerFinish)
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

func (ui *UI) createNotificationButton() *widget.Button {
	var b *widget.Button
	b = widget.NewButtonWithIcon("", theme.VolumeUpIcon(), func() {
		sound.toggle()

		if sound.enabled {
			b.Icon = theme.VolumeUpIcon()
		} else {
			b.Icon = theme.VolumeMuteIcon()
		}

		b.Refresh()
	})

	s := theme.IconInlineSize() * float32(NotificationButtonMultiplier)

	b.Move(fyne.NewPos(WindowWidth-s, 0))
	b.Resize(fyne.NewSize(s, s))

	return b
}

func (ui *UI) disableBreaks() {
	for i := range ui.breaks {
		ui.breaks[i].disable()
	}
}
