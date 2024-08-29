package instructions

import (
	"errors"
)

type valueBuilder struct {
	reference Reference
	bytes     []byte
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		reference: nil,
		bytes:     nil,
	}

	return &out
}

// Create initializes the valueBuilder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithReference adds a reference to the valueBuilder
func (app *valueBuilder) WithReference(reference Reference) ValueBuilder {
	app.reference = reference
	return app
}

// WithBytes add bytes to the valueBuilder
func (app *valueBuilder) WithBytes(bytes []byte) ValueBuilder {
	app.bytes = bytes
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes != nil {
		return createValueWithBytes(app.bytes), nil
	}

	if app.reference != nil {
		return createValueWithReference(app.reference), nil
	}

	return nil, errors.New("the Value is invalid")
}
