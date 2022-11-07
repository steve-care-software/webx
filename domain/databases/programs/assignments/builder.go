package assignments

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	pIndex *uint
	value  entities.Identifier
}

func createBuilder() Builder {
	out := builder{
		pIndex: nil,
		value:  nil,
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
func (app *builder) WithValue(value entities.Identifier) Builder {
	app.value = value
	return app
}

// Now builds a new Assignment instance
func (app *builder) Now() (Assignment, error) {
	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Assignment instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Assignment instance")
	}

	return createAssignment(*app.pIndex, app.value), nil
}
