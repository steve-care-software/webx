package cmain

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type customEntry struct {
	widget.Entry
}

func newCustomEntry() *customEntry {
	entry := &customEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (c *customEntry) CreateRenderer() fyne.WidgetRenderer {
	// Create a custom background
	background := canvas.NewRectangle(color.RGBA{R: 255, B: 203, G: 192, A: 255})

	// Use the default renderer for the entry
	entryRenderer := c.Entry.CreateRenderer()

	// Replace the default background with the custom one
	objects := entryRenderer.Objects()
	objects[0] = background

	return &customEntryRenderer{
		WidgetRenderer: entryRenderer,
		background:     background,
	}
}

type customEntryRenderer struct {
	fyne.WidgetRenderer
	background *canvas.Rectangle
}

func (r *customEntryRenderer) Refresh() {
	r.background.FillColor = color.RGBA{R: 20, B: 20, G: 20, A: 255}
	r.background.StrokeColor = color.RGBA{R: 255, B: 203, G: 192, A: 255}
	r.background.StrokeWidth = 1
	r.background.StrokeColor = theme.Color(theme.ColorNameBackground)
	r.background.SetMinSize(fyne.NewSize(200, 50))
	r.background.CornerRadius = 15

	r.background.Refresh()

	// Call the base renderer's refresh to apply other styles
	r.WidgetRenderer.Refresh()
}

type mainIns struct {
	application fyne.App
	window      fyne.Window
}

func createMain(
	application fyne.App,
	window fyne.Window,
) Main {
	out := mainIns{
		application: application,
		window:      window,
	}

	return &out
}

// Fetch fetches the container
func (app *mainIns) Fetch() *fyne.Container {

	titleWidget := newCustomEntry()
	titleWidget.SetText("myContainer")
	/*titleWidget.TextStyle = fyne.TextStyle{
		Bold:      true,
		Italic:    true,
		Underline: true,
	}*/

	/*color := color.RGBA{R: 20, B: 20, G: 20, A: 255}
	rect := canvas.NewRectangle(color)
	rect.StrokeWidth = 1
	rect.StrokeColor = theme.Color(theme.ColorNameBackground)
	rect.SetMinSize(fyne.NewSize(200, 50))
	rect.CornerRadius = 15

	titleStack := container.NewStack(
		titleWidget,
		rect,
	)*/

	/*blackColor := color.RGBA{R: 125, B: 125, G: 125, A: 255}
	rect := canvas.NewRectangle(blackColor)
	rect.SetMinSize(fyne.NewSize(50, 100))

	titleBox := container.NewHBox(
		widget.NewLabel("myBlock"),
		layout.NewSpacer(),
		widget.NewButton("Save", func() {
			log.Printf("save...")
		}),
		widget.NewButton("Copy", func() {
			log.Printf("copy...")
		}),
		widget.NewButton("Minimize", func() {
			log.Printf("minimize...")
		}),
		widget.NewButton("Delete", func() {
			log.Printf("delete...")
		}),
	)

	paddedContainer := container.NewPadded(
		rect,
		titleBox,
	)*/

	/*blueColor := color.RGBA{R: 0, G: 128, B: 255, A: 255}
	rect := canvas.NewRectangle(blueColor) // Blue color
	rect.SetMinSize(fyne.NewSize(200, 100))

	popupContent := container.NewVBox(
		widget.NewLabel("myBlock"),
	)

	popup := widget.NewPopUp(popupContent, app.window.Canvas())

	popupContent.Objects = append(
		popupContent.Objects,
		rect,
		widget.NewButton("Close", func() {
			popup.Hide()
		}),
	)

	// Create the button that will trigger the popup
	triggerButton := widget.NewButton("myBlock", nil)
	triggerButton.OnTapped = func() {
		// Create a popup and position it relative to the triggerButton
		popup.ShowAtPosition(fyne.NewPos(triggerButton.Position().X+50, triggerButton.Position().Y+triggerButton.Size().Height+50))
	}*/

	containerIns := container.New(layout.NewVBoxLayout(), titleWidget)

	// Create a container with a max layout to fill 100% width and height
	textFieldContainer := container.New(layout.NewStackLayout(), containerIns)

	// Create a container with the widget filling the central space
	containerWithFullWidget := container.NewBorder(
		nil, nil, nil, nil, // No top, bottom, left, or right widgets
		textFieldContainer, // Center widget fills the available space
	)

	// Create a red rectangle to use as the background
	//redBackground := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})

	// Create a container with the red background and the content on top
	//return container.NewStack(redBackground, textFieldContainer, redBackground)

	return containerWithFullWidget
}
