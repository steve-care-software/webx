package grids

import (
	"fyne.io/fyne/v2"
	center_center "github.com/steve-care-software/webx/gui/grids/center"
	center_left "github.com/steve-care-software/webx/gui/grids/left"
	center_right "github.com/steve-care-software/webx/gui/grids/right"
	center_top "github.com/steve-care-software/webx/gui/grids/top"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	topBuilder := center_top.NewBuilder()
	leftBuilder := center_left.NewBuilder()
	centerBuilder := center_center.NewBuilder()
	rightBuilder := center_right.NewBuilder()
	return createBuilder(
		topBuilder,
		leftBuilder,
		centerBuilder,
		rightBuilder,
	)
}

// Builder represents the center builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Grid, error)
}

// Grid represents the center container in the grid
type Grid interface {
	Fetch() *fyne.Container
}
