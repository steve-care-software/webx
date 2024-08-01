package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
)

type pointerBuilder struct {
	delimiter delimiters.Delimiter
	isDeleted bool
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		delimiter: nil,
		isDeleted: false,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithDelimiter adds a delimiter to the builder
func (app *pointerBuilder) WithDelimiter(delimiter delimiters.Delimiter) PointerBuilder {
	app.delimiter = delimiter
	return app
}

// IsDeleted flags the builder as deleted
func (app *pointerBuilder) IsDeleted() PointerBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.delimiter == nil {
		return nil, errors.New("the delimiter is mandatory in order to build a Pointer instance")
	}

	if app.isDeleted {
		return createPointerWithDeleted(app.delimiter), nil
	}

	return createPointer(app.delimiter), nil
}
