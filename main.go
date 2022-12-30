package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

var (
	ui       *UI
	breaks   []*BreakWidget
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

func main() {
	sound = new(Sound)
	sound.initContext()
	sound.initCache()
	sound.cache["notification"] = sound.open(NotificationSound)

	ui = new(UI)
	ui.app = app.New()

	ui.app.Settings().SetTheme(&newTheme{})

	ui.window = ui.app.NewWindow(WindowTitle)

	if desk, ok := ui.app.(desktop.App); ok {
		m := fyne.NewMenu(WindowTitle,
			fyne.NewMenuItem("Show", func() {
				ui.window.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}

	ui.window.SetContent(ui.createContent())
	ui.window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	ui.window.SetCloseIntercept(func() { ui.window.Hide() })
	ui.window.ShowAndRun()
}
