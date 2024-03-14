package origins

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

type builder struct {
	hashAdapter hash.Adapter
	resource    resources.Resource
	operator    operators.Operator
	next        Value
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		resource:    nil,
		operator:    nil,
		next:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource builder
func (app *builder) WithResource(resource resources.Resource) Builder {
	app.resource = resource
	return app
}

// WithOperator adds an operator builder
func (app *builder) WithOperator(operator operators.Operator) Builder {
	app.operator = operator
	return app
}

// WithNext adds a next value to the builder builder
func (app *builder) WithNext(next Value) Builder {
	app.next = next
	return app
}

// Now builds a new Origin instance
func (app *builder) Now() (Origin, error) {
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
