package references

import "errors"

type pointerBuilder struct {
	pFrom  *uint
	length uint
}

func createPointerBuilder() PointerBuilder {
	out := pointerBuilder{
		pFrom:  nil,
		length: 0,
	}

	return &out
}

// Create initializes the builder
func (app *pointerBuilder) Create() PointerBuilder {
	return createPointerBuilder()
}

// WithLength adds a length to the builder
func (app *pointerBuilder) WithLength(length uint) PointerBuilder {
	app.length = length
	return app
}

// From adds a from to the builder
func (app *pointerBuilder) From(from uint) PointerBuilder {
	app.pFrom = &from
	return app
}

// Now builds a new Pointer instance
func (app *pointerBuilder) Now() (Pointer, error) {
	if app.pFrom == nil {
		return nil, errors.New("the from index is mandatory in order to build a Pointer instance")
	}

	if app.length <= 0 {
		return nil, errors.New("the length must be greater than zero (0) in order to build a Pointer instance")
	}

	return createPointer(*app.pFrom, app.length), nil
}
