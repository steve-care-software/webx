package left

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type left struct {
	application fyne.App
	window      fyne.Window
}

func createLeft(
	application fyne.App,
	window fyne.Window,
) Left {
	out := left{
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the container
func (app *left) Fetch() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Left Container"),
	)
}
