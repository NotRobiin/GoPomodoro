package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type newTheme struct{}

var _ fyne.Theme = (*newTheme)(nil)

func (m newTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	none := color.RGBA{R: 0, G: 0, B: 0, A: 0}

	switch name {
	case theme.ColorNameButton, theme.ColorNameHover:
		return none
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m newTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m newTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m newTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNamePadding {
		return 0
	}

	return theme.DefaultTheme().Size(name)
}
