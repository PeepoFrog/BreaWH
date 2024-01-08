package gui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/PeepoFrog/breawg/dialogs"
)

func (g *Gui) StartBreathing() {
	var wizard *dialogs.Wizard
	someContent := widget.NewButton("stop", func() {
		// g.mainContent.Objects = g.mainContent.Objects[0:1]
		wizard.Hide()
	})
	// g.mainContent.Objects[2] = someContent
	wizard = dialogs.NewWizard("t", someContent)

	wizard.Show(g.Window)
	wizard.Resize(g.mainContent.Size())
}
