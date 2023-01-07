package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SettingsWidget struct {
	toggleButton *widget.Button
	overlay      *widget.PopUp
	widget       *fyne.Container
	enabled      bool
	title        string
}

func (sw *SettingsWidget) create(c fyne.Canvas, onDismiss func()) {
	b := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() { sw.toggle() })
	b.Move(fyne.NewPos(WindowWidth-theme.IconInlineSize(), 0))
	b.Resize(fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize()))

	title := container.New(layout.NewCenterLayout(), canvas.NewText(sw.title, color.White))
	con := container.New(layout.NewMaxLayout(), title, container.New(layout.NewHBoxLayout(), layout.NewSpacer(), b))

	sw.widget = container.New(layout.NewVBoxLayout(), con, widget.NewSeparator())
	sw.overlay = widget.NewModalPopUp(sw.widget, c)
}

func (sw *SettingsWidget) add(title string, obj fyne.CanvasObject, layoutType fyne.Layout, spacer bool) {
	t := canvas.NewText(title, color.White)
	c := container.New(layout.NewMaxLayout(), obj)
	con := container.New(layoutType)
	con.Add(t)

	if spacer {
		con.Add(layout.NewSpacer())
	}

	con.Add(c)

	sw.widget.Add(con)
}

func (sw *SettingsWidget) toggle() {
	if sw.enabled {
		sw.hide()
	} else {
		sw.show()
	}

	sw.enabled = !sw.enabled
}

func (sw *SettingsWidget) hide() {
	sw.widget.Hide()
	sw.overlay.Hide()
}

func (sw *SettingsWidget) show() {
	sw.widget.Show()
	sw.overlay.Show()
}
