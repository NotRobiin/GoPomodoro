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
	overlayShown          bool
}

func NewTapPopUp(content fyne.CanvasObject, canvas fyne.Canvas, onTap, onTapSecondary func()) *TapPopUp {
	return &TapPopUp{PopUp: widget.NewPopUp(content, canvas), onTap: onTap, onTapSecondary: onTapSecondary}
}

func (p *TapPopUp) Tapped(_ *fyne.PointEvent) {
	p.onTap()
	p.PopUp.Tapped(nil)
}

func (p *TapPopUp) TappedSecondary(_ *fyne.PointEvent) {
	p.onTapSecondary()
	p.PopUp.TappedSecondary(nil)
}

func (p *TapPopUp) Show() {
	if !p.overlayShown {
		p.Canvas.Overlays().Add(p)
		p.overlayShown = true
	}
	p.Refresh()
	p.BaseWidget.Show()
}

func (p *TapPopUp) Hide() {
	if p.overlayShown {
		p.Canvas.Overlays().Remove(p)
		p.overlayShown = false
	}
	p.BaseWidget.Hide()
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