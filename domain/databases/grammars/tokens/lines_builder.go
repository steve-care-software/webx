package tokens

import "errors"

type linesBuilder struct {
	list []Line
}

func createLinesBuilder() LinesBuilder {
	out := linesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *linesBuilder) Create() LinesBuilder {
	return createLinesBuilder()
}

// WithList adds a list to the builder
func (app *linesBuilder) WithList(list []Line) LinesBuilder {
	app.list = list
	return app
}

// Now builds a new Lines instance
func (app *linesBuilder) Now() (Lines, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Line in order to build a Lines instance")
	}

	return createLines(app.list), nil
}
