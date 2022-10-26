package selections

import "errors"

type selectionBuilder struct {
	elementName string
	list        []Child
}

func createSelectionBuilder() SelectionBuilder {
	out := selectionBuilder{
		elementName: "",
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *selectionBuilder) Create() SelectionBuilder {
	return createSelectionBuilder()
}

// WithElementName adds an element name to the builder
func (app *selectionBuilder) WithElementName(elementName string) SelectionBuilder {
	app.elementName = elementName
	return app
}

// WithList adds a list to the builder
func (app *selectionBuilder) WithList(list []Child) SelectionBuilder {
	app.list = list
	return app
}

// Now builds a new Selection instance
func (app *selectionBuilder) Now() (Selection, error) {
	if app.elementName == "" {
		return nil, errors.New("the element name is mandatory in order to build a Selection instance")
	}

	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Child in order to build a Selection instance")
	}

	return createSelection(app.elementName, app.list), nil
}
