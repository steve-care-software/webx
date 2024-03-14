package resources

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	layer       hash.Hash
	layerBytes  []byte
	isMandatory bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		layer:       nil,
		layerBytes:  nil,
		isMandatory: false,
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
func (app *builder) WithLayer(layer hash.Hash) Builder {
	app.layer = layer
	return app
}

// IsMandatory flags the builder as mandatory
func (app *builder) IsMandatory() Builder {
	app.isMandatory = true
	return app
}

// WithLayerBytes add layer bytes to the builder
func (app *builder) WithLayerBytes(layerBytes []byte) Builder {
	app.layerBytes = layerBytes
	return app
}

// Now builds a new Resource instance
func (app *builder) Now() (Resource, error) {
	if app.layerBytes != nil && len(app.layerBytes) <= 0 {
		app.layerBytes = nil
	}

	if app.layerBytes != nil {
		pHash, err := app.hashAdapter.FromBytes(app.layerBytes)
		if err != nil {
			return nil, err
		}

		app.layer = *pHash
	}

	if app.layer == nil {
		return nil, errors.New("the layer hash is mandatory in order to build an Resource instance")
	}

	isMandatory := "false"
	if app.isMandatory {
		isMandatory = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.layer.Bytes(),
		[]byte(isMandatory),
	})

	if err != nil {
		return nil, err
	}

	return createResource(*pHash, app.layer, app.isMandatory), nil
}
