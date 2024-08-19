package asts

import (
	"errors"
)

type builder struct {
	list []AST
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []AST) Builder {
	app.list = list
	return app
}

// Now builds a new ASTs instance
func (app *builder) Now() (ASTs, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 AST in order to build a ASTs instance")
	}

	return createASTs(app.list), nil
}
