package databases

import (
	"errors"
	"time"
)

type pointerBuilder struct {
	beginsOn   SizeInBytes
	length     SizeInBytes
	pCreatedOn *time.Time
	references Pointers
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		beginsOn:   nil,
		length:     nil,
		pCreatedOn: nil,
		references: nil,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithLength adds a length to the builder
func (app *pointerBuilder) WithLength(length SizeInBytes) PointerBuilder {
	app.length = length
	return app
}

// WithReferences add references to the builder
func (app *pointerBuilder) WithReferences(references Pointers) PointerBuilder {
	app.references = references
	return app
}

// BeginsOn adds a beginsOn size to the builder
func (app *pointerBuilder) BeginsOn(beginsOn SizeInBytes) PointerBuilder {
	app.beginsOn = beginsOn
	return app
}

// CreatedOn adds a creation time to the builder
func (app *pointerBuilder) CreatedOn(createdOn time.Time) PointerBuilder {
	app.pCreatedOn = &createdOn
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.length == nil {
		return nil, errors.New("the length is mandatory in order to build a Pointer instance")
	}

	if app.beginsOn == nil {
		return nil, errors.New("the beginsOn is mandatory in order to build a Pointer instance")
	}

	if app.pCreatedOn == nil {
		return nil, errors.New("the creation time is mandatory in order to build a Pointer instance")
	}

	if app.length.IsZero() {
		return nil, errors.New("the length must be greater than zero (0)")
	}

	if app.references != nil {
		return createPointerWithReferences(app.beginsOn, app.length, *app.pCreatedOn, app.references), nil
	}

	return createPointer(app.beginsOn, app.length, *app.pCreatedOn), nil
}
