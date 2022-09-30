package uncovers

import "errors"

type lineBuilder struct {
	pIndex   *uint
	elements []string
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		pIndex:   nil,
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithIndex adds an index to the builder
func (app *lineBuilder) WithIndex(index uint) LineBuilder {
	app.pIndex = &index
	return app
}

// WithElements add elements to the builder
func (app *lineBuilder) WithElements(elements []string) LineBuilder {
	app.elements = elements
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Line instance")
	}

	if app.elements != nil && len(app.elements) <= 0 {
		app.elements = nil
	}

	if app.elements == nil {
		return nil, errors.New("thre must be at least 1 element in order to build a Line instance")
	}

	return createLine(*app.pIndex, app.elements), nil
}
