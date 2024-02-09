package links

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type originValueBuilder struct {
	hashAdapter hash.Adapter
	resource    OriginResource
	origin      Origin
}

func createOriginValueBuilder(
	hashAdapter hash.Adapter,
) OriginValueBuilder {
	out := originValueBuilder{
		hashAdapter: hashAdapter,
		resource:    nil,
		origin:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *originValueBuilder) Create() OriginValueBuilder {
	return createOriginValueBuilder(
		app.hashAdapter,
	)
}

// WithResource adds a resource to the builder
func (app *originValueBuilder) WithResource(resource OriginResource) OriginValueBuilder {
	app.resource = resource
	return app
}

// WithOrigin adds an origin to the builder
func (app *originValueBuilder) WithOrigin(origin Origin) OriginValueBuilder {
	app.origin = origin
	return app
}

// Now builds a new OriginValue instance
func (app *originValueBuilder) Now() (OriginValue, error) {
	data := [][]byte{}
	if app.resource != nil {
		data = append(data, app.resource.Hash().Bytes())
	}

	if app.origin != nil {
		data = append(data, app.origin.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the OriginValue is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.resource != nil {
		return createOriginValueWithResource(*pHash, app.resource), nil
	}

	return createOriginValueWithOrigin(*pHash, app.origin), nil
}
