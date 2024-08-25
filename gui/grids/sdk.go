package grids

import (
	"fyne.io/fyne/v2"
	"github.com/steve-care-software/webx/gui/grids/bottom"
	"github.com/steve-care-software/webx/gui/grids/center"
	"github.com/steve-care-software/webx/gui/grids/top"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	bottomBuilder := bottom.NewBuilder()
	centerBuilder := center.NewBuilder()
	topBuilder := top.NewBuilder()
	return createBuilder(
		bottomBuilder,
		centerBuilder,
		topBuilder,
	)
}

// Builder represents the grid builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Grid, error)
}

// Grid represents the application grid
type Grid interface {
	Fetch() *fyne.Container
}
