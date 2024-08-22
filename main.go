package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/steve-care-software/webx/gui/menus"
)

func main() {
	application := app.New()
	window := application.NewWindow("Steve Care")
	mainMenu, err := menus.NewMainMenuBuilder().Create().
		WithApplication(application).
		WithWindow(window).
		Now()

	if err != nil {
		panic(err)
	}

	window.SetMainMenu(mainMenu.Fetch())
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
