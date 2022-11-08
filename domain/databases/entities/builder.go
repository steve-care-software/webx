package entities

import "errors"

type builder struct {
	identifier Identifier
	signature  Signature
}

func createBuilder() Builder {
	out := builder{
		identifier: nil,
		signature:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier Identifier) Builder {
	app.identifier = identifier
	return app
}

// WithSignature adds a signature to the builder
func (app *builder) WithSignature(signature Signature) Builder {
	app.signature = signature
	return app
}

// Now builds a new Entity instance
func (app *builder) Now() (Entity, error) {
	if app.identifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build an Entity instance")
	}

	if app.signature == nil {
		return nil, errors.New("the signature is mandatory in order to build an Entity instance")
	}

	return createEntity(app.identifier, app.signature), nil
}
