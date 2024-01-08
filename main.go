package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/PeepoFrog/breawg/gui"
)

func main() {
	fmt.Println("Pepega")
	a := app.NewWithID("BreaWG")

	gui := gui.Gui{}
	gui.Window = a.NewWindow("BreaWH")

	gui.Window.SetMaster()
	gui.Window.Resize(fyne.NewSize(400, 800))
	content := gui.MakeGui()

	gui.Window.SetContent(content)
	gui.Window.ShowAndRun()
}
