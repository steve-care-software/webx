package elements

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
)

type elementBuilder struct {
	hashAdapter hash.Adapter
	layer       hash.Hash
	layerBytes  []byte
	condition   conditions.Condition
}

func createElementBuilder(
	hashAdapter hash.Adapter,
) ElementBuilder {
	out := elementBuilder{
		hashAdapter: hashAdapter,
		layer:       nil,
		layerBytes:  nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder(
		app.hashAdapter,
	)
}

// WithLayer adds a layer to the builder
func (app *elementBuilder) WithLayer(layer hash.Hash) ElementBuilder {
	app.layer = layer
	return app
}

// WithCondition adds a condition to the builder
func (app *elementBuilder) WithCondition(condition conditions.Condition) ElementBuilder {
	app.condition = condition
	return app
}

// WithLayerBytes adds layer bytes to the builder
func (app *elementBuilder) WithLayerBytes(layerBytes []byte) ElementBuilder {
	app.layerBytes = layerBytes
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
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
		return nil, errors.New("the layer hash is mandatory in order to build an Element instance")
	}

	data := [][]byte{
		app.layer.Bytes(),
	}

	if app.condition != nil {
		data = append(data, app.condition.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.condition != nil {
		return createElementWithCondition(*pHash, app.layer, app.condition), nil
	}

	return createElement(*pHash, app.layer), nil
}
