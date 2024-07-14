package amounts

import "errors"

type builder struct {
	context string
	ret     string
}

func createBuilder() Builder {
	out := builder{
		context: "",
		ret:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// WithReturn adds a return to the builder
func (app *builder) WithReturn(ret string) Builder {
	app.ret = ret
	return app
}

// Now builds a new Amount instance
func (app *builder) Now() (Amount, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build an Amount instance")
	}

	if app.ret == "" {
		return nil, errors.New("the return is mandatory in order to build an Amount instance")
	}

	return createAmount(app.context, app.ret), nil
}
