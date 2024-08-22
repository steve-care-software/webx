package menus

import (
	"errors"

	"fyne.io/fyne/v2"
)

type mainMenuBuilder struct {
	application fyne.App
	window      fyne.Window
}

func createMainMenuBuilder() MainMenuBuilder {
	out := mainMenuBuilder{
		application: nil,
		window:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *mainMenuBuilder) Create() MainMenuBuilder {
	return createMainMenuBuilder()
}

// WithApplication adds an application to the builder
func (app *mainMenuBuilder) WithApplication(application fyne.App) MainMenuBuilder {
	app.application = application
	return app
}

// WithWindow adds a window to the builder
func (app *mainMenuBuilder) WithWindow(window fyne.Window) MainMenuBuilder {
	app.window = window
	return app
}

// Now builds a new MainMenu instance
func (app *mainMenuBuilder) Now() (MainMenu, error) {
	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build a MainMenu instance")
	}

	if app.window == nil {
		return nil, errors.New("the window is mandatory in order to build a MainMenu instance")
	}

	return createMainMenu(
		app.application,
		app.window,
	), nil
}
