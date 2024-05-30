package references

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

type referenceBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	instance    instances.Instance
}

func createReferenceBuilder(
	hashAdapter hash.Adapter,
) ReferenceBuilder {
	out := referenceBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		instance:    nil,
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

// WithInstance adds an instance to the builder
func (app *referenceBuilder) WithInstance(instance instances.Instance) ReferenceBuilder {
	app.instance = instance
	return app
}

// Now builds a new Reference instance
func (app *referenceBuilder) Now() (Reference, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Reference instance")
	}

	if app.instance == nil {
		return nil, errors.New("the instance is mandatory in order to build a Reference instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.variable),
		app.instance.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createReference(*pHash, app.variable, app.instance), nil
}
