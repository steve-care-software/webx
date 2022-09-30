package grammars

import "errors"

type everythingBuilder struct {
	exception Line
	escape    Line
}

func createEverythingBuilder() EverythingBuilder {
	out := everythingBuilder{
		exception: nil,
		escape:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *everythingBuilder) Create() EverythingBuilder {
	return createEverythingBuilder()
}

// WithException adds an exception to the builder
func (app *everythingBuilder) WithException(exception Line) EverythingBuilder {
	app.exception = exception
	return app
}

// WithEscape adds an escape to the builder
func (app *everythingBuilder) WithEscape(escape Line) EverythingBuilder {
	app.escape = escape
	return app
}

// Now builds a new Everything instance
func (app *everythingBuilder) Now() (Everything, error) {
	if app.exception == nil {
		return nil, errors.New("the exception is mandatory in order to build an Everything instance")
	}

	if app.escape != nil {
		return createEverythingWithEscape(app.exception, app.escape), nil
	}

	return createEverything(app.exception), nil
}
