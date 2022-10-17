package indexes

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Index
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
	return createBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Index) Builder {
	app.list = list
	return app
}

// Now builds a new Indexes instance
func (app *builder) Now() (Indexes, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Index in order to build an Indexes instance")
	}

	data := [][]byte{}
	for _, oneIndex := range app.list {
		data = append(data, oneIndex.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createIndexes(*pHash, app.list), nil
}
