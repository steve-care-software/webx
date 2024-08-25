package gui

import (
	"errors"

	fine_app "fyne.io/fyne/v2/app"
	"github.com/steve-care-software/webx/gui/grids"
	"github.com/steve-care-software/webx/gui/menus"
)

type builder struct {
	gridBuilder grids.Builder
	menuBuilder menus.Builder
	title       string
	width       float32
	height      float32
}

func createBuilder(
	gridBuilder grids.Builder,
	menuBuilder menus.Builder,
) Builder {
	out := builder{
		gridBuilder: gridBuilder,
		menuBuilder: menuBuilder,
		title:       "",
		width:       0,
		height:      0,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.gridBuilder,
		app.menuBuilder,
	)
}

// WithTitle adds a title to the builder
func (app *builder) WithTitle(title string) Builder {
	app.title = title
	return app
}

// WithWidth adds a width to the builder
func (app *builder) WithWidth(width float32) Builder {
	app.width = width
	return app
}

// WithHeight adds an height to the builder
func (app *builder) WithHeight(height float32) Builder {
	app.height = height
	return app
}

// Now builds a new Gui instance
func (app *builder) Now() (Gui, error) {
	if app.title == "" {
		return nil, errors.New("the title is mandatory in order to build a Gui")
	}

	if app.width <= 0 {
		return nil, errors.New("the width is mandatory in order to build a Gui")
	}

	if app.height <= 0 {
		return nil, errors.New("the height is mandatory in order to build a Gui")
	}

	application := fine_app.New()
	window := application.NewWindow(app.title)
	grid, err := app.gridBuilder.Create().
		WithApplication(application).
		WithWindow(window).
		Now()

	if err != nil {
		return nil, err
	}

	menu, err := app.menuBuilder.
		WithApplication(application).
		WithWindow(window).
		Now()

	if err != nil {
		return nil, err
	}

	return createGui(
		app.title,
		app.width,
		app.height,
		grid,
		menu,
		application,
		window,
	), nil
}
