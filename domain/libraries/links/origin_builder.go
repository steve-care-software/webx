package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type originBuilder struct {
	hashAdapter hash.Adapter
	resource    OriginResource
	operator    Operator
	next        OriginValue
}

func createOriginBuilder(
	hashAdapter hash.Adapter,
) OriginBuilder {
	out := originBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		operator:    nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *originBuilder) Create() OriginBuilder {
	return createOriginBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource builder
func (app *originBuilder) WithResource(resource OriginResource) OriginBuilder {
	app.resource = resource
	return app
}

// WithOperator adds an operator builder
func (app *originBuilder) WithOperator(operator Operator) OriginBuilder {
	app.operator = operator
	return app
}

// WithNext adds a next value to the builder builder
func (app *originBuilder) WithNext(next OriginValue) OriginBuilder {
	app.next = next
	return app
}

// Now builds a new Origin instance
func (app *originBuilder) Now() (Origin, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Origin instance")
	}

	if app.operator == nil {
		return nil, errors.New("the operator is mandatory in order to build an Origin instance")
	}

	if app.next == nil {
		return nil, errors.New("the next value is mandatory in order to build an Origin instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.resource.Hash().Bytes(),
		app.operator.Hash().Bytes(),
		app.next.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createOrigin(*pHash, app.resource, app.operator, app.next), nil
}
