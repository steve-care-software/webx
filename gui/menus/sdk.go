package menus

import "fyne.io/fyne/v2"

// NewMainMenuBuilder creates a new main menu builder
func NewMainMenuBuilder() MainMenuBuilder {
	return createMainMenuBuilder()
}

// MainMenuBuilder creates the main menu builder
type MainMenuBuilder interface {
	Create() MainMenuBuilder
	WithApplication(application fyne.App) MainMenuBuilder
	WithWindow(window fyne.Window) MainMenuBuilder
	Now() (MainMenu, error)
}

// MainMenu creates the main menu
type MainMenu interface {
	Fetch() *fyne.MainMenu
}
