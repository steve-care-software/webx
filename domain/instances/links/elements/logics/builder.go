package logics

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics/locations"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

type builder struct {
	hashAdapter hash.Adapter
	layer       layers.Layer
	location    locations.Location
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layer:       nil,
		location:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer layers.Layer) Builder {
	app.layer = layer
	return app
}

// WithLocation adds a location to the builder
func (app *builder) WithLocation(location locations.Location) Builder {
	app.location = location
	return app
}

// Now builds a new Logic instance
func (app *builder) Now() (Logic, error) {
	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build a Logic instance")
	}

	if app.location == nil {
		return nil, errors.New("the location is mandatory in order to build a Logic instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.layer.Hash().Bytes(),
		app.location.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLogic(*pHash, app.layer, app.location), nil
}
