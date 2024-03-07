package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type originResourceBuilder struct {
	hashAdapter hash.Adapter
	layer       hash.Hash
	layerBytes  []byte
	isMandatory bool
}

func createOriginResourceBuilder(
	hashAdapter hash.Adapter,
) OriginResourceBuilder {
	out := originResourceBuilder{
		hashAdapter: hashAdapter,
		layer:       nil,
		layerBytes:  nil,
		isMandatory: false,
	}

	return &out
}

// Create initializes the builder
func (app *originResourceBuilder) Create() OriginResourceBuilder {
	return createOriginResourceBuilder(
		app.hashAdapter,
	)
}

// WithLayer adds a layer to the builder
func (app *originResourceBuilder) WithLayer(layer hash.Hash) OriginResourceBuilder {
	app.layer = layer
	return app
}

// IsMandatory flags the builder as mandatory
func (app *originResourceBuilder) IsMandatory() OriginResourceBuilder {
	app.isMandatory = true
	return app
}

// WithLayerBytes add layer bytes to the builder
func (app *originResourceBuilder) WithLayerBytes(layerBytes []byte) OriginResourceBuilder {
	app.layerBytes = layerBytes
	return app
}

// Now builds a new OriginResource instance
func (app *originResourceBuilder) Now() (OriginResource, error) {
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
		return nil, errors.New("the layer hash is mandatory in order to build an OriginResource instance")
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

	return createOriginResource(*pHash, app.layer, app.isMandatory), nil
}
