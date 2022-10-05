package grammars

import "errors"

type everythingBuilder struct {
	name      string
	exception Token
	escape    Token
}

func createEverythingBuilder() EverythingBuilder {
	out := everythingBuilder{
		name:      "",
		exception: nil,
		escape:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *everythingBuilder) Create() EverythingBuilder {
	return createEverythingBuilder()
}

// WithName adds a name to the builder
func (app *everythingBuilder) WithName(name string) EverythingBuilder {
	app.name = name
	return app
}

// WithException adds an exception to the builder
func (app *everythingBuilder) WithException(exception Token) EverythingBuilder {
	app.exception = exception
	return app
}

// WithEscape adds an escape to the builder
func (app *everythingBuilder) WithEscape(escape Token) EverythingBuilder {
	app.escape = escape
	return app
}

// Now builds a new Everything instance
func (app *everythingBuilder) Now() (Everything, error) {
	if app.exception == nil {
		return nil, errors.New("the exception is mandatory in order to build an Everything instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Everything instance")
	}

	if app.escape != nil {
		return createEverythingWithEscape(app.name, app.exception, app.escape), nil
	}

	return createEverything(app.name, app.exception), nil
}
