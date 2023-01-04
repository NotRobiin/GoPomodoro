package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type TapPopUp struct {
	*widget.PopUp

	onTap, onTapSecondary func()
}

func NewTapPopUp(content fyne.CanvasObject, canvas fyne.Canvas, onTap, onTapSecondary func()) *TapPopUp {
	return &TapPopUp{widget.NewPopUp(content, canvas), onTap, onTapSecondary}
}

func (p *TapPopUp) Tapped(_ *fyne.PointEvent) {
	p.onTap()
	p.PopUp.Tapped(nil)
}

func (p *TapPopUp) TappedSecondary(_ *fyne.PointEvent) {
	p.onTapSecondary()
	p.PopUp.TappedSecondary(nil)
}

type SettingsWidget struct {
	toggleButton *widget.Button
	overlay      *TapPopUp
	widget       *fyne.Container
	enabled      bool
}

func (sw *SettingsWidget) create(c fyne.Canvas, onDismiss func(), settings ...fyne.CanvasObject) {
	sw.widget = container.New(layout.NewVBoxLayout(), settings...)
	sw.overlay = NewTapPopUp(sw.widget, c, onDismiss, onDismiss)
}

func (sw *SettingsWidget) toggle() {
	sw.enabled = !sw.enabled

	if sw.enabled {
		sw.show()
	} else {
		sw.hide()
	}
}

func (sw *SettingsWidget) hide() {
	sw.widget.Hide()
	sw.overlay.Hide()
}

func (sw *SettingsWidget) show() {
	sw.widget.Show()
	sw.overlay.Show()
}
