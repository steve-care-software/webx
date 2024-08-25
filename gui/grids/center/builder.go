package center

import (
	"errors"

	"fyne.io/fyne/v2"
	"github.com/steve-care-software/webx/gui/grids/center/bottom"
	"github.com/steve-care-software/webx/gui/grids/center/cmain"
)

type builder struct {
	mainBuilder   cmain.Builder
	bottomBuilder bottom.Builder
	application   fyne.App
	window        fyne.Window
}

func createBuilder(
	mainBuilder cmain.Builder,
	bottomBuilder bottom.Builder,
) Builder {
	out := builder{
		mainBuilder:   mainBuilder,
		bottomBuilder: bottomBuilder,
		application:   nil,
		window:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.mainBuilder,
		app.bottomBuilder,
	)
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

// Now builds a new Center instance
func (app *builder) Now() (Center, error) {
	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build a Center instance")
	}

	if app.window == nil {
		return nil, errors.New("the window is mandatory in order to build a Center instance")
	}

	mainIns, err := app.mainBuilder.Create().WithApplication(app.application).WithWindow(app.window).Now()
	if err != nil {
		return nil, err
	}

	bottom, err := app.bottomBuilder.Create().WithApplication(app.application).WithWindow(app.window).Now()
	if err != nil {
		return nil, err
	}

	return createCenter(
		mainIns,
		bottom,
		app.application,
		app.window,
	), nil
}
