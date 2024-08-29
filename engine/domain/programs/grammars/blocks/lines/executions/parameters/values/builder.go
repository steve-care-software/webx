package values

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values/references"
)

type builder struct {
	reference references.Reference
	bytes     []byte
}

func createBuilder() Builder {
	out := builder{
		reference: nil,
		bytes:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithReference adds a reference to the builder
func (app *builder) WithReference(reference references.Reference) Builder {
	app.reference = reference
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
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
