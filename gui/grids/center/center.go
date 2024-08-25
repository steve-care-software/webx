package center

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/steve-care-software/webx/gui/grids/center/bottom"
	"github.com/steve-care-software/webx/gui/grids/center/cmain"
)

type center struct {
	main        cmain.Main
	bottom      bottom.Bottom
	application fyne.App
	window      fyne.Window
}

func createCenter(
	main cmain.Main,
	bottom bottom.Bottom,
	application fyne.App,
	window fyne.Window,
) Center {
	out := center{
		main:        main,
		bottom:      bottom,
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the container
func (app *center) Fetch() *fyne.Container {
	mainContainer := app.main.Fetch()
	bottomContainer := app.bottom.Fetch()
	return container.NewBorder(
		nil,             // Top
		bottomContainer, // Bottom
		nil,             // Left
		nil,             // Right
		mainContainer,   // Grid (expands)
	)
}
