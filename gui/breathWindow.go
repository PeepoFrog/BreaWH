package gui

import "fyne.io/fyne/v2/widget"

func (g *Gui) StartBreathing() {
	someContent := widget.NewButton("stop", func() {
		g.mainContent.Objects = g.mainContent.Objects[0:1]
	})
	g.mainContent.Objects = append(g.mainContent.Objects, someContent)
}
