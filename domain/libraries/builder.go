package libraries

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/links"
)

type builder struct {
	hashAdapter hash.Adapter
	layers      layers.Layers
	links       links.Links
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layers:      nil,
		links:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLayers add layers to the builder
func (app *builder) WithLayers(layers layers.Layers) Builder {
	app.layers = layers
	return app
}

// WithLinks add links to the builder
func (app *builder) WithLinks(links links.Links) Builder {
	app.links = links
	return app
}

// Now builds a new Library instance
func (app *builder) Now() (Library, error) {
	if app.layers == nil {
		return nil, errors.New("the layers is mandatory in order to build a Library instance")
	}

	data := [][]byte{
		app.layers.Hash().Bytes(),
	}

	if app.links != nil {
		data = append(data, app.links.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.links != nil {
		return createLibraryWithLinks(*pHash, app.layers, app.links), nil
	}

	return createLibrary(*pHash, app.layers), nil
}
