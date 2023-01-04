package results

import "errors"

type builder struct {
	isValid   bool
	values    []interface{}
	remaining []byte
}

func createBuilder() Builder {
	out := builder{
		isValid:   false,
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
func (app *builder) WithValues(values []interface{}) Builder {
	app.values = values
	return app
}

// WithRemaining adds a remaining to the builder
func (app *builder) WithRemaining(remaining []byte) Builder {
	app.remaining = remaining
	return app
}

// IsValid flags the builder as valid
func (app *builder) IsValid() Builder {
	app.isValid = true
	return app
}

// Now builds a new Result instance
func (app *builder) Now() (Result, error) {
	if app.values != nil && len(app.values) <= 0 {
		app.values = nil
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = nil
	}

	if !app.isValid && app.values != nil {
		return nil, errors.New("the Result cannot contain values because it is invalid")
	}

	if app.values != nil && app.remaining != nil {
		return createResultWithValuesAndRemaining(app.isValid, app.values, app.remaining), nil
	}

	if app.values != nil {
		return createResultWithValues(app.isValid, app.values), nil
	}

	if app.remaining != nil {
		return createResultWithRemaining(app.isValid, app.remaining), nil
	}

	return createResult(app.isValid), nil
}
