package top

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type top struct {
	application fyne.App
	window      fyne.Window
}

func createTop(
	application fyne.App,
	window fyne.Window,
) Top {
	out := top{
		application: application,
		window:      window,
	}
	return &out
}

// Fetch fetches the top container
func (app *top) Fetch() *fyne.Container {
	text4 := canvas.NewText("top", color.White)
	contentContainer := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		text4,
		layout.NewSpacer(),
	)

	// Create a blue to serve as the background
	blueContainer := canvas.NewRectangle(color.RGBA{R: 0, G: 0, B: 255, A: 255})

	// Set the stack:
	return container.NewStack(blueContainer, contentContainer)
}
