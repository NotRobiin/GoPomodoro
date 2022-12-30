package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var ui *UI
var breaks []*BreakWidget

func formatTime(tm time.Duration) string {
	s := int(tm.Seconds())
	minutes := int(s/60) % 60
	seconds := int(s % 60)

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func main() {
	ui = new(UI)
	ui.app = app.New()

	ui.app.Settings().SetTheme(&newTheme{})

	ui.window = ui.app.NewWindow(WindowTitle)
	ui.window.SetContent(ui.createContent())
	ui.window.Resize(fyne.NewSize(WindowWidth, WindowHeight))
	ui.window.SetMaster()
	ui.window.CenterOnScreen()
	ui.window.RequestFocus()
	ui.window.SetFixedSize(true)
	ui.window.ShowAndRun()
}
