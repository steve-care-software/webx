package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/elements/conditions/resources"
)

type builder struct {
	hashAdapter hash.Adapter
	resource    resources.Resource
	next        ConditionValue
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		resource:    nil,
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

// WithNext adds a next value to the builder builder
func (app *builder) WithNext(next ConditionValue) Builder {
	app.next = next
	return app
}

// Now builds a new Condition instance
func (app *builder) Now() (Condition, error) {
	if app.resource == nil {
		return nil, errors.New("the resource is mandatory in order to build an Condition instance")
	}

	data := [][]byte{
		app.resource.Hash().Bytes(),
	}
	if app.next != nil {
		data = append(data, app.next.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.next != nil {
		return createConditionWithNext(*pHash, app.resource, app.next), nil
	}

	return createCondition(*pHash, app.resource), nil
}
