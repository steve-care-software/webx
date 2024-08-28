package cmain

import (
	"errors"

	"fyne.io/fyne/v2"
)

type builder struct {
	application fyne.App
	window      fyne.Window
}

func createBuilder() Builder {
	out := builder{
		application: nil,
		window:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithApplication adds an application to the builder
func (app *builder) WithApplication(application fyne.App) Builder {
	app.application = application
	return app
}

// WithWindow adds a window to the builder
func (app *builder) WithWindow(window fyne.Window) Builder {
	app.window = window
	return app
}

// Now builds a new Main instance
func (app *builder) Now() (Main, error) {
	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build a Center instance")
	}

	if app.window == nil {
		return nil, errors.New("the window is mandatory in order to build a Center instance")
	}

	return createMain(
		app.application,
		app.window,
	), nil
}