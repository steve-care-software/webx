package center

import (
	"fyne.io/fyne/v2"
	"github.com/steve-care-software/webx/gui/grids/center/bottom"
	"github.com/steve-care-software/webx/gui/grids/center/cmain"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	mainBuilder := cmain.NewBuilder()
	bottomBuilder := bottom.NewBuilder()
	return createBuilder(
		mainBuilder,
		bottomBuilder,
	)
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Center, error)
}

// Center represents the center container
type Center interface {
	Fetch() *fyne.Container
}
