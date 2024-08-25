package grids

import (
	"errors"

	"fyne.io/fyne/v2"
	"github.com/steve-care-software/webx/gui/grids/bottom"
	"github.com/steve-care-software/webx/gui/grids/center"
	"github.com/steve-care-software/webx/gui/grids/top"
)

type builder struct {
	bottomBuilder bottom.Builder
	centerBuilder center.Builder
	topBuilder    top.Builder
	application   fyne.App
	window        fyne.Window
}

func createBuilder(
	bottomBuilder bottom.Builder,
	centerBuilder center.Builder,
	topBuilder top.Builder,
) Builder {
	out := builder{
		bottomBuilder: bottomBuilder,
		centerBuilder: centerBuilder,
		topBuilder:    topBuilder,
		application:   nil,
		window:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.bottomBuilder,
		app.centerBuilder,
		app.topBuilder,
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

// Now builds a new Grid instance
func (app *builder) Now() (Grid, error) {
	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build a Grid instance")
	}

	if app.window == nil {
		return nil, errors.New("the window is mandatory in order to build a Grid instance")
	}

	bottom, err := app.bottomBuilder.Create().
		WithApplication(app.application).
		WithWindow(app.window).
		Now()

	if err != nil {
		return nil, err
	}

	center, err := app.centerBuilder.Create().
		WithApplication(app.application).
		WithWindow(app.window).
		Now()

	if err != nil {
		return nil, err
	}

	top, err := app.topBuilder.Create().
		WithApplication(app.application).
		WithWindow(app.window).
		Now()

	if err != nil {
		return nil, err
	}

	return createGrid(
		bottom,
		center,
		top,
	), nil
}
