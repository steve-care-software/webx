package top

import "fyne.io/fyne/v2"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the top builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Top, error)
}

// Top represents the top container in the grid
type Top interface {
	Fetch() *fyne.Container
}
