package bridges

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

type bridgeBuilder struct {
	hashAdapter hash.Adapter
	path        []string
	layer       layers.Layer
}

func createBridgeBuilder(
	hashAdapter hash.Adapter,
) BridgeBuilder {
	out := bridgeBuilder{
		hashAdapter: hashAdapter,
		path:        nil,
		layer:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *bridgeBuilder) Create() BridgeBuilder {
	return createBridgeBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *bridgeBuilder) WithPath(path []string) BridgeBuilder {
	app.path = path
	return app
}

// WithLayer adds a layer to the builder
func (app *bridgeBuilder) WithLayer(layer layers.Layer) BridgeBuilder {
	app.layer = layer
	return app
}

// Now builds a new Bridge instance
func (app *bridgeBuilder) Now() (Bridge, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Bridge instance")
	}

	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build a Bridge instance")
	}

	path := filepath.Join(app.path...)
	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(path),
		app.layer.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createBridge(*pHash, app.path, app.layer), nil
}
