package outputs

import (
	"errors"
)

type builder struct {
	values    map[string]interface{}
	remaining []byte
}

func createBuilder() Builder {
	out := builder{
		values:    nil,
		remaining: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithValues add values to the builder
func (app *builder) WithValues(values map[string]interface{}) Builder {
	app.values = values
	return app
}

// WithRemaining adds a remaining to the builder
func (app *builder) WithRemaining(remaining []byte) Builder {
	app.remaining = remaining
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
	if app.values == nil {
		return nil, errors.New("the values is mandatory in order to build an Output instance")
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = nil
	}

	if app.remaining != nil {
		return createOutputWithRemaining(app.values, app.remaining), nil
	}

	return createOutput(app.values), nil
}
