package grids

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/steve-care-software/webx/gui/grids/bottom"
	"github.com/steve-care-software/webx/gui/grids/center"
	"github.com/steve-care-software/webx/gui/grids/top"
)

type grid struct {
	bottom bottom.Bottom
	center center.Center
	top    top.Top
}

func createGrid(
	bottom bottom.Bottom,
	center center.Center,
	top top.Top,
) Grid {
	out := grid{
		bottom: bottom,
		center: center,
		top:    top,
	}

	return &out
}

// Fetch fetches the grid container
func (app *grid) Fetch() *fyne.Container {
	top := app.top.Fetch()
	center := app.center.Fetch()
	bottom := app.bottom.Fetch()
	return container.New(
		layout.NewVBoxLayout(),
		top,
		center,
		bottom,
	)
}
