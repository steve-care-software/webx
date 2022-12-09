package assignments

import (
	"errors"

	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

type builder struct {
	pIndex *uint
	pValue *hash.Hash
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		pValue: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// WithValue adds a value to the builder
func (app *builder) WithValue(value hash.Hash) Builder {
	app.pValue = &value
	return app
}

// Now builds a new Assignment instance
func (app *builder) Now() (Assignment, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Assignment instance")
	}

	if app.pValue == nil {
		return nil, errors.New("the value is mandatory in order to build an Assignment instance")
	}

	return createAssignment(*app.pIndex, *app.pValue), nil
}
