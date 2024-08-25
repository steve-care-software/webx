package left

import "fyne.io/fyne/v2"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Left, error)
}

// Left represents the left container
type Left interface {
	Fetch() *fyne.Container
}
