package layers

import "github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers/pointers"

type layerBuilder struct {
	isDeleted bool
	pointers  pointers.Pointers
}

func createLayerBuilder() LayerBuilder {
	out := layerBuilder{
		isDeleted: false,
		pointers:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *layerBuilder) Create() LayerBuilder {
	return createLayerBuilder()
}

// WithPointers adds pointers to the builder
func (app *layerBuilder) WithPointers(pointers pointers.Pointers) LayerBuilder {
	app.pointers = pointers
	return app
}

// IsDeleted flags the builder as deleted
func (app *layerBuilder) IsDeleted() LayerBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new Layer instance
func (app *layerBuilder) Now() (Layer, error) {
	if app.pointers != nil {
		return createLayerWithPointers(app.isDeleted, app.pointers), nil
	}

	return createLayer(app.isDeleted), nil
}
