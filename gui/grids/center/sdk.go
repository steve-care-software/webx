package center

import "fyne.io/fyne/v2"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the center builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Center, error)
}

// Center represents the center container in the grid
type Center interface {
	Fetch() *fyne.Container
}
