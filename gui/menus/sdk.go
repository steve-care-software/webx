package menus

import "fyne.io/fyne/v2"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder creates the main menu builder
type Builder interface {
	Create() Builder
	WithApplication(application fyne.App) Builder
	WithWindow(window fyne.Window) Builder
	Now() (MainMenu, error)
}

// MainMenu creates the main menu
type MainMenu interface {
	Fetch() *fyne.MainMenu
}
