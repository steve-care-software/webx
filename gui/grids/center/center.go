package center

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type center struct {
	application fyne.App
	window      fyne.Window
}

func createCenter(
	application fyne.App,
	window fyne.Window,
) Center {
	out := center{
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the the container
func (app *center) Fetch() *fyne.Container {
	text4 := canvas.NewText("center", color.White)
	contentContainer := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		text4,
		layout.NewSpacer(),
	)

	// Create a green to serve as the background
	greenContainer := canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 255})

	// Set the stack:
	return container.NewStack(greenContainer, contentContainer)
}
