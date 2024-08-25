package grids

import (
	"errors"

	"fyne.io/fyne/v2"
	center_center "github.com/steve-care-software/webx/gui/grids/center"
	center_left "github.com/steve-care-software/webx/gui/grids/left"
	center_right "github.com/steve-care-software/webx/gui/grids/right"
	center_top "github.com/steve-care-software/webx/gui/grids/top"
)

type builder struct {
	topBuilder    center_top.Builder
	leftBuilder   center_left.Builder
	centerBuilder center_center.Builder
	rightBuilder  center_right.Builder
	application   fyne.App
	window        fyne.Window
}

func createBuilder(
	topBuilder center_top.Builder,
	leftBuilder center_left.Builder,
	centerBuilder center_center.Builder,
	rightBuilder center_right.Builder,
) Builder {
	out := builder{
		topBuilder:    topBuilder,
		leftBuilder:   leftBuilder,
		centerBuilder: centerBuilder,
		rightBuilder:  rightBuilder,
		application:   nil,
		window:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.topBuilder,
		app.leftBuilder,
		app.centerBuilder,
		app.rightBuilder,
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

// Now builds a new Bottom instance
func (app *builder) Now() (Grid, error) {
	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build a Grid instance")
	}

	if app.window == nil {
		return nil, errors.New("the window is mandatory in order to build a Grid instance")
	}

	top, err := app.topBuilder.Create().WithApplication(app.application).WithWindow(app.window).Now()
	if err != nil {
		return nil, err
	}

	left, err := app.leftBuilder.Create().WithApplication(app.application).WithWindow(app.window).Now()
	if err != nil {
		return nil, err
	}

	center, err := app.centerBuilder.Create().WithApplication(app.application).WithWindow(app.window).Now()
	if err != nil {
		return nil, err
	}

	right, err := app.rightBuilder.Create().WithApplication(app.application).WithWindow(app.window).Now()
	if err != nil {
		return nil, err
	}

	return createGrid(
		top,
		left,
		center,
		right,
		app.application,
		app.window,
	), nil
}
