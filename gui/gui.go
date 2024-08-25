package gui

import (
	"fyne.io/fyne/v2"
	"github.com/steve-care-software/webx/gui/grids"
	"github.com/steve-care-software/webx/gui/menus"
)

type gui struct {
	title       string
	width       float32
	height      float32
	grid        grids.Grid
	mainMenu    menus.MainMenu
	application fyne.App
	window      fyne.Window
}

func createGui(
	title string,
	width float32,
	height float32,
	grid grids.Grid,
	mainMenu menus.MainMenu,
	application fyne.App,
	window fyne.Window,
) Gui {
	out := gui{
		title:       title,
		width:       width,
		height:      height,
		grid:        grid,
		mainMenu:    mainMenu,
		application: application,
		window:      window,
	}

	return &out
}

// Execute executes the gui
func (app *gui) Execute() error {
	// set the content:
	gridContainer := app.grid.Fetch()
	app.window.SetContent(gridContainer)

	// set the menu:
	mainMenu := app.mainMenu.Fetch()
	app.window.SetMainMenu(mainMenu)

	// resize the window:
	app.window.Resize(fyne.NewSize(app.width, app.height))

	// show and run the window:
	app.window.ShowAndRun()
	return nil
}
