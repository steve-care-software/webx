package bottom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	text4 := canvas.NewText("bottom", color.White)
	contentContainer := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		text4,
		layout.NewSpacer(),
	)

	// Create a red rectangle to serve as the background
	redBackground := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Set the stack:
	return container.NewStack(redBackground, contentContainer)
}
