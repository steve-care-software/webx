package bridges

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Bridge
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
func (app *builder) WithList(list []Bridge) Builder {
	app.list = list
	return app
}

// Now builds a new Bridges instance
func (app *builder) Now() (Bridges, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Bridge in order to build a Bridges instance")
	}

	data := [][]byte{}
	mp := map[string]Bridge{}
	for _, oneBridge := range app.list {
		path := filepath.Join(oneBridge.Path()...)
		mp[path] = oneBridge
		data = append(data, oneBridge.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createBridges(*pHash, app.list, mp), nil

}
