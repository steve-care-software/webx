package right

import "fyne.io/fyne/v2"

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Right, error)
}

// Right represents the right container
type Right interface {
	Fetch() (*fyne.Container, error)
}
