package pointers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
)

type pointerBuilder struct {
	retrieval retrievals.Retrieval
	isDeleted bool
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		retrieval: nil,
		isDeleted: false,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithRetrieval adds a retrieval to the builder
func (app *pointerBuilder) WithRetrieval(retrieval retrievals.Retrieval) PointerBuilder {
	app.retrieval = retrieval
	return app
}

// IsDeleted flags the builder as deleted
func (app *pointerBuilder) IsDeleted() PointerBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.retrieval == nil {
		return nil, errors.New("the retrieval is mandatory in order to build a Pointer instance")
	}

	if app.isDeleted {
		return createPointerWithDeleted(app.retrieval), nil
	}

	return createPointer(app.retrieval), nil
}
