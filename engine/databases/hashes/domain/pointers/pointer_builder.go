package pointers

import (
	"errors"

	bytes_pointers "github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type pointerBuilder struct {
	hash    hash.Hash
	pointer bytes_pointers.Pointer
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		hash:    nil,
		pointer: nil,
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

// WithPointer adds a pointer to the builder
func (app *pointerBuilder) WithPointer(pointer bytes_pointers.Pointer) PointerBuilder {
	app.pointer = pointer
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.hash == nil {
		return nil, errors.New("the hash is mandatory in order to build a Pointer instance")
	}

	if app.pointer == nil {
		return nil, errors.New("the pointer is mandatory in order to build a Pointer instance")
	}

	return createPointer(
		app.hash,
		app.pointer,
	), nil
}
