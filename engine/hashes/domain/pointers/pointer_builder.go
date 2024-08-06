package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
)

type pointerBuilder struct {
	hash      hash.Hash
	delimiter delimiters.Delimiter
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		hash:      nil,
		delimiter: nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithHash adds an hash to the builder
func (app *pointerBuilder) WithHash(hash hash.Hash) PointerBuilder {
	app.hash = hash
	return app
}

// WithDelimiter adds a delimiter to the builder
func (app *pointerBuilder) WithDelimiter(delimiter delimiters.Delimiter) PointerBuilder {
	app.delimiter = delimiter
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Pointer instance")
	}

	if app.delimiter == nil {
		return nil, errors.New("the delimiter is mandatory in order to build a Pointer instance")
	}

	return createPointer(
		app.hash,
		app.delimiter,
	), nil
}
