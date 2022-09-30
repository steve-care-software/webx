package coverages

import "errors"

type lineBuilder struct {
	list []Element
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithList adds a list to the builder
func (app *lineBuilder) WithList(list []Element) LineBuilder {
	app.list = list
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Element in order to build a Line instance")
	}

	return createLine(app.list), nil
}
