package references

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Reference
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
func (app *builder) WithList(list []Reference) Builder {
	app.list = list
	return app
}

// Now builds a new References instance
func (app *builder) Now() (References, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Reference in order to build a References instance")
	}

	data := [][]byte{}
	for _, oneReference := range app.list {
		data = append(data, oneReference.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mp := map[string]Reference{}
	for _, oneReference := range app.list {
		keyname := oneReference.Hash().String()
		mp[keyname] = oneReference
	}

	return createReferences(*pHash, mp, app.list), nil
}
