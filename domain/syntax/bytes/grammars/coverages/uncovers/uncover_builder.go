package uncovers

import "errors"

type uncoverBuilder struct {
	name string
	line Line
}

func createUncoverBuilder() UncoverBuilder {
	out := uncoverBuilder{
		name: "",
		line: nil,
	}

	return &out
}

// Create initializes the builder
func (app *uncoverBuilder) Create() UncoverBuilder {
	return createUncoverBuilder()
}

// WithName adds a name to the builder
func (app *uncoverBuilder) WithName(name string) UncoverBuilder {
	app.name = name
	return app
}

// WithLine adds a line to the builder
func (app *uncoverBuilder) WithLine(line Line) UncoverBuilder {
	app.line = line
	return app
}

// WithLine adds a line to the builder
func (app *uncoverBuilder) Now() (Uncover, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Uncover instance")
	}

	if app.line == nil {
		return nil, errors.New("the line is mandatory in order to build an Uncover instance")
	}

	return createUncover(app.name, app.line), nil
}
