package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Layer
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Layer) Builder {
	app.list = list
	return app
}

// Now builds a new Layers instance
func (app *builder) Now() (Layers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Layer in order to build a Layers instance")
	}

	data := [][]byte{}
	for _, oneLayer := range app.list {
		data = append(data, oneLayer.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mp := map[string]Layer{}
	for _, oneLayer := range app.list {
		keyname := oneLayer.Hash().String()
		mp[keyname] = oneLayer
	}

	return createLayers(*pHash, mp, app.list), nil
}
