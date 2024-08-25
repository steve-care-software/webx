package bottom

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type bottom struct {
	application fyne.App
	window      fyne.Window
}

func createBottom(
	application fyne.App,
	window fyne.Window,
) Bottom {
	out := bottom{
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the the container
func (app *bottom) Fetch() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Bottom Container"),
	)
}
