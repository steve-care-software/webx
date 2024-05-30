package references

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type referenceBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	identifier  hash.Hash
}

func createReferenceBuilder(
	hashAdapter hash.Adapter,
) ReferenceBuilder {
	out := referenceBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		identifier:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *referenceBuilder) Create() ReferenceBuilder {
	return createReferenceBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *referenceBuilder) WithVariable(variable string) ReferenceBuilder {
	app.variable = variable
	return app
}

// WithIdentifier adds an identifier to the builder
func (app *referenceBuilder) WithIdentifier(identifier hash.Hash) ReferenceBuilder {
	app.identifier = identifier
	return app
}

// Now builds a new Reference instance
func (app *referenceBuilder) Now() (Reference, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Reference instance")
	}

	if app.identifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Reference instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		app.identifier.Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createReference(*pHash, app.variable, app.identifier), nil
}
