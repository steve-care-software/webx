package logics

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

type logicBuilder struct {
	hashAdapter hash.Adapter
	link        links.Link
	layers      layers.Layers
}

func createLogicBuilder(
	hashAdapter hash.Adapter,
) LogicBuilder {
	out := logicBuilder{
		hashAdapter: hashAdapter,
		link:        nil,
		layers:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *logicBuilder) Create() LogicBuilder {
	return createLogicBuilder(
		app.hashAdapter,
	)
}

// WithLink adds a link to the builder
func (app *logicBuilder) WithLink(link links.Link) LogicBuilder {
	app.link = link
	return app
}

// WithLayers add layers to the builder
func (app *logicBuilder) WithLayers(layers layers.Layers) LogicBuilder {
	app.layers = layers
	return app
}

// Now builds a new Logic instance
func (app *logicBuilder) Now() (Logic, error) {
	if app.link == nil {
		return nil, errors.New("the link is mandatory in order to build a Logic instance")
	}

	if app.layers == nil {
		return nil, errors.New("the layers is mandatory in order to build a Logic instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.link.Hash().Bytes(),
		app.layers.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLogic(*pHash, app.link, app.layers), nil
}
