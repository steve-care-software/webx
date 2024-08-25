package grids

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	center_center "github.com/steve-care-software/webx/gui/grids/center"
	center_left "github.com/steve-care-software/webx/gui/grids/left"
	center_right "github.com/steve-care-software/webx/gui/grids/right"
	center_top "github.com/steve-care-software/webx/gui/grids/top"
)

type center struct {
	top         center_top.Top
	left        center_left.Left
	center      center_center.Center
	right       center_right.Right
	application fyne.App
	window      fyne.Window
}

func createGrid(
	top center_top.Top,
	left center_left.Left,
	centerIns center_center.Center,
	right center_right.Right,
	application fyne.App,
	window fyne.Window,
) Grid {
	out := center{
		top:         top,
		left:        left,
		center:      centerIns,
		right:       right,
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the the container
func (app *center) Fetch() *fyne.Container {
	topContainer := app.top.Fetch()
	leftContainer := app.left.Fetch()
	centerContainer := app.center.Fetch()
	rightContainer := app.right.Fetch()
	return container.NewBorder(
		topContainer,    // Top
		nil,             // Bottom
		leftContainer,   // Left
		rightContainer,  // Right
		centerContainer, // Grid (expands)
	)
}
