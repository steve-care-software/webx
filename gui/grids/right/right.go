package right

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type right struct {
	application fyne.App
	window      fyne.Window
}

func createRight(
	application fyne.App,
	window fyne.Window,
) Right {
	out := right{
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the container
func (app *right) Fetch() *fyne.Container {
	return container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Right Container"),
	)
}
