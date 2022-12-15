package contents

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type valueBuilder struct {
	pByte *byte
	tree  entities.Identifier
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		pByte: nil,
		tree:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithByte adds byte to the builder
func (app *valueBuilder) WithByte(byte byte) ValueBuilder {
	app.pByte = &byte
	return app
}

// WithTree adds a tree to the builder
func (app *valueBuilder) WithTree(tree entities.Identifier) ValueBuilder {
	app.tree = tree
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.pByte != nil {
		return createValueWithByte(app.pByte), nil
	}

	if app.tree != nil {
		return createValueWithTree(app.tree), nil
	}

	return nil, errors.New("the Value is invalid")
}
