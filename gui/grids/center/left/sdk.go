package left

import "fyne.io/fyne/v2"

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Left, error)
}

// Left represents the left container
type Left interface {
	Fetch() (*fyne.Container, error)
}
