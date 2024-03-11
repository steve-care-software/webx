package libraries

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/references"
)

type builder struct {
	hashAdapter hash.Adapter
	layers      layers.Layers
	links       links.Links
	references  references.References
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layers:      nil,
		links:       nil,
		references:  nil,
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

// WithReferences add references to the builder
func (app *builder) WithReferences(references references.References) Builder {
	app.references = references
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

	if app.references != nil {
		data = append(data, app.references.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.references != nil && app.links != nil {
		return createLibraryWithLinksAndReferences(*pHash, app.layers, app.links, app.references), nil
	}

	if app.links != nil {
		return createLibraryWithLinks(*pHash, app.layers, app.links), nil
	}

	if app.references != nil {
		return createLibraryWithReferences(*pHash, app.layers, app.references), nil
	}

	return createLibrary(*pHash, app.layers), nil
}
