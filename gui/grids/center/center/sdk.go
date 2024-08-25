package center

import "fyne.io/fyne/v2"

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (Center, error)
}

// Center represents the center container
type Center interface {
	Fetch() (*fyne.Container, error)
}
