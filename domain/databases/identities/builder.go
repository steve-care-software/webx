package identities

import (
	"errors"

	"github.com/steve-care-software/webx/domain/databases/entities"
)

type builder struct {
	identifier    entities.Identifier
	modifications entities.Identifiers
}

func createBuilder() Builder {
	out := builder{
		identifier:    nil,
		modifications: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier entities.Identifier) Builder {
	app.identifier = identifier
	return app
}

// WithModifications add modifications to the builder
func (app *builder) WithModifications(modifications entities.Identifiers) Builder {
	app.modifications = modifications
	return app
}

// Now builds a new Identity instance
func (app *builder) Now() (Identity, error) {
	if app.identifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build an Identity instance")
	}

	if app.modifications == nil {
		return nil, errors.New("the modifications is mandatory in order to build an Identity instance")
	}

	return createIdentity(app.identifier, app.modifications), nil
}
