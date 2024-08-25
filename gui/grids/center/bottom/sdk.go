package bottom

import "fyne.io/fyne/v2"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the bottom builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Bottom, error)
}

// Bottom represents the bottom container in the grid
type Bottom interface {
	Fetch() *fyne.Container
}
