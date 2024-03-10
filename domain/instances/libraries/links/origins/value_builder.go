package origins

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/origins/resources"
)

type valueBuilder struct {
	hashAdapter hash.Adapter
	resource    resources.Resource
	origin      Origin
}

func createValueBuilder(
	hashAdapter hash.Adapter,
) ValueBuilder {
	out := valueBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		origin:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource to the builder
func (app *valueBuilder) WithResource(resource resources.Resource) ValueBuilder {
	app.resource = resource
	return app
}

// WithOrigin adds an origin to the builder
func (app *valueBuilder) WithOrigin(origin Origin) ValueBuilder {
	app.origin = origin
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	data := [][]byte{}
	if app.resource != nil {
		data = append(data, app.resource.Hash().Bytes())
	}

	if app.origin != nil {
		data = append(data, app.origin.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Value is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.resource != nil {
		return createValueWithResource(*pHash, app.resource), nil
	}

	return createValueWithOrigin(*pHash, app.origin), nil
}
