package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

type builder struct {
	hashAdapter hash.Adapter
	input       []byte
	source      links.Link
	layers      layers.Layers
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		input:       nil,
		source:      nil,
		layers:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input []byte) Builder {
	app.input = input
	return app
}

// WithSource adds a source to the builder
func (app *builder) WithSource(source links.Link) Builder {
	app.source = source
	return app
}

// WithLayers adds a layers to the builder
func (app *builder) WithLayers(layers layers.Layers) Builder {
	app.layers = layers
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Link instance")
	}

	if app.source == nil {
		return nil, errors.New("the source is mandatory in order to build a Link instance")
	}

	data := [][]byte{
		app.input,
		app.source.Hash().Bytes(),
	}

	if app.layers != nil {
		data = append(data, app.layers.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.layers != nil {
		return createLinkWithLayers(*pHash, app.input, app.source, app.layers), nil
	}

	return createLink(*pHash, app.input, app.source), nil
}
