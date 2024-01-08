package gui

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/widget"
)

var (
	repeatsCounter   = binding.NewInt()
	infCheck         = binding.NewBool()
	minutesToHold    = binding.NewInt()
	secondsToHold    = binding.NewInt()
	secondsToAddjust = binding.NewInt() // how much seconds will be added each round

	hold time.Time
)

type Gui struct {
	Window      fyne.Window
	mainContent *fyne.Container
}

func setDefaultTimer() {
	minutesToHold.Set(1)
	secondsToHold.Set(0)
	secondsToAddjust.Set(0)

	err := repeatsCounter.Set(3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func (g *Gui) MakeGui() fyne.CanvasObject {
	setDefaultTimer()
	g.mainContent = container.NewStack()
	g.mainContent.Objects = []fyne.CanvasObject{settingsMenu()} //settingsMenu()

	content := container.NewBorder(nil,
		widget.NewButton("START", func() {
			g.StartBreathing()
		}), nil, nil,

		g.mainContent,
	)
	// content := container.NewVBox(
	// 	makeBreatheHoldCnahgerRow(),

	// 	makeRepeatsRow(),
	// 	container.NewStack(),
	// 	widget.NewButton("START", nil),
	// )

	return content
}

func settingsMenu() fyne.CanvasObject {
	menu := container.NewVBox(
		makeBreatheHoldCnahgerRow(),
		makeRepeatsRow(),
	)
	return menu
}

func minutesToHoldChanger() fyne.CanvasObject {
	minutesToHold.Set(1)

	minutesToHoldLabel := widget.NewLabel("1")
	minutesToHoldLabel.Alignment = 1
	changeMinutes := func(numToChange int) {

		i, err := minutesToHold.Get()
		fmt.Println(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		i += numToChange
		if i > 0 && i <= 60 {
			minutesToHold.Set(i)
			minutesToHoldLabel.SetText(strconv.Itoa(i))
		}
		// minutesToHold.Set(numToChange)

	}
	minutesToHoldChanger := container.NewVBox(
		widget.NewButton(" +1 ", func() { changeMinutes(1) }),
		minutesToHoldLabel,
		widget.NewButton(" -1 ", func() { changeMinutes(-1) }),
	)
	return minutesToHoldChanger
}

func secondsToHoldChanger() fyne.CanvasObject {
	// secondsToHoldLabel := widget.NewLabel("0")
	s, err := secondsToAddjust.Get()
	if err != nil {
		s = 0
	}
	secondsToHoldLabel := widget.NewLabel(strconv.Itoa(s))
	secondsToHoldLabel.Alignment = 1

	changeSeconds := func(numToChange int) {
		i, err := secondsToHold.Get()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("secondsToHold.Get()", i)
		i += numToChange
		if i >= 0 && i <= 60 {
			secondsToHold.Set(i)
			fmt.Println("secondsToHold.Set()", i)
			fmt.Println("")
			secondsToHoldLabel.SetText(strconv.Itoa(i))
		}
	}
	secondsToHoldChanger := container.NewVBox(
		widget.NewButton(" +5 ", func() { changeSeconds(5) }),
		secondsToHoldLabel,
		widget.NewButton(" -5 ", func() { changeSeconds(-5) }),
	)
	return secondsToHoldChanger
}

func secondsToHoldAddjusterChanger() fyne.CanvasObject {
	s, err := secondsToAddjust.Get()
	if err != nil {
		s = 0
	}
	secondsToAddjustLabel := widget.NewLabel(strconv.Itoa(s))
	secondsToAddjustLabel.Alignment = 1

	changeSecondsToAddjust := func(numToChange int) {
		i, err := secondsToAddjust.Get()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("secondsToHold.Get()", i)
		i += numToChange
		if i >= 0 && i <= 60 {
			secondsToAddjust.Set(i)
			fmt.Println("changeSecondsToAddjust()", i)
			fmt.Println("")
			secondsToAddjustLabel.SetText(strconv.Itoa(i))
		}
	}

	secondsToHoldChanger := container.NewVBox(
		widget.NewButton(" +1 ", func() { changeSecondsToAddjust(1) }),
		secondsToAddjustLabel,
		widget.NewButton(" -1 ", func() { changeSecondsToAddjust(-1) }),
	)
	return secondsToHoldChanger
}

func makeBreatheHoldCnahgerRow() fyne.CanvasObject {
	//minutes changer to hold

	// seconds changer to hold
	infoText := widget.NewLabel("+ seconds each round")
	infoText.Wrapping = 2
	timerEdditor := container.NewGridWithColumns(4,
		minutesToHoldChanger(),
		secondsToHoldChanger(),
		infoText,
		secondsToHoldAddjusterChanger(),
	)
	return timerEdditor
}

func makeRepeatsRow() fyne.CanvasObject {
	n, err := repeatsCounter.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	counterLabel := widget.NewLabel(strconv.Itoa(n))
	counterLabel.Alignment = 1
	changeCounterLabel := func(numToChange int) {
		i, err := repeatsCounter.Get()
		if err != nil {
			fmt.Println(err)
		}
		i += numToChange
		if i != 0 {
			repeatsCounter.Set(i)

			counterLabel.SetText(strconv.Itoa(i))

		}
		fmt.Println(i)
	}

	repeatsRow := container.NewVBox(container.NewGridWithRows(3,
		widget.NewButton(" +1 ", func() { changeCounterLabel(+1) }),
		counterLabel,
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton(" -1 ", func() { changeCounterLabel(-1) }),
			layout.NewSpacer(),
		),
	),
		widget.NewCheck("inf...", func(b bool) {
			infCheck.Set(b)
			if b {
				counterLabel.SetText("inf...")
			} else {
				i, err := repeatsCounter.Get()
				if err != nil {
					repeatsCounter.Set(3)
				}
				counterLabel.SetText(strconv.Itoa(i))

			}
		}),
	)
	return repeatsRow
}

func makeSettingTab() fyne.CanvasObject {
	return container.NewStack()
}
