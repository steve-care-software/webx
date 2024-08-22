package menus

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type mainMenu struct {
	application fyne.App
	window      fyne.Window
}

func createMainMenu(
	application fyne.App,
	window fyne.Window,
) MainMenu {
	out := mainMenu{
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the main menu
func (app *mainMenu) Fetch() *fyne.MainMenu {
	// Create menu items
	newItem := fyne.NewMenuItem("New", func() {
		// Action for "New"
	})

	openItem := fyne.NewMenuItem("Open", func() {
		// Action for "Open"
	})

	saveItem := fyne.NewMenuItem("Save", func() {
		// Action for "Save"
	})

	quitItem := fyne.NewMenuItem("Quit", func() {
		app.application.Quit()
	})

	// Create a submenu under File
	fileMenu := fyne.NewMenu("File", newItem, openItem, saveItem, quitItem)

	// Create an Edit menu
	cutItem := fyne.NewMenuItem("Cut", func() {
		// Action for "Cut"
	})

	copyItem := fyne.NewMenuItem("Copy", func() {
		// Action for "Copy"
	})

	pasteItem := fyne.NewMenuItem("Paste", func() {
		// Action for "Paste"
	})

	editMenu := fyne.NewMenu("Edit", cutItem, copyItem, pasteItem)

	// Create a Help menu
	aboutItem := fyne.NewMenuItem("About", func() {
		dialog := widget.NewLabel("This is an example application.")
		app.window.SetContent(container.NewCenter(dialog))
	})

	helpMenu := fyne.NewMenu("Help", aboutItem)

	// Create the main menu and attach it to the window
	return fyne.NewMainMenu(fileMenu, editMenu, helpMenu)
}
