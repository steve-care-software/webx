package selections

import (
	"errors"
)

type selectionBuilder struct {
	element  Element
	children Children
}

func createSelectionBuilder() SelectionBuilder {
	out := selectionBuilder{
		element:  nil,
		children: nil,
	}

	return &out
}

// Create initializes the builder
func (app *selectionBuilder) Create() SelectionBuilder {
	return createSelectionBuilder()
}

// WithElement adds an element to the builder
func (app *selectionBuilder) WithElement(element Element) SelectionBuilder {
	app.element = element
	return app
}

// WithChildren adds a children to the builder
func (app *selectionBuilder) WithChildren(children Children) SelectionBuilder {
	app.children = children
	return app
}

// Now builds a new Selection instance
func (app *selectionBuilder) Now() (Selection, error) {
	if app.element != nil && app.children != nil {
		content := createContentWithElementAndChildren(app.element, app.children)
		return createSelection(content), nil
	}

	if app.element != nil {
		content := createContentWithElement(app.element)
		return createSelection(content), nil
	}

	if app.children != nil {
		content := createContentWithChildren(app.children)
		return createSelection(content), nil
	}

	return nil, errors.New("the Selection is invalid")
}
