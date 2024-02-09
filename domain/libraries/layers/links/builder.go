package links

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Link
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
func (app *builder) WithList(list []Link) Builder {
	app.list = list
	return app
}

// Now builds a new Links instance
func (app *builder) Now() (Links, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Link in order to build a Links instance")
	}

	data := [][]byte{}
	for _, oneLink := range app.list {
		data = append(data, oneLink.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	mp := map[string]Link{}
	for _, oneLink := range app.list {
		keyname := oneLink.Hash().String()
		mp[keyname] = oneLink
	}

	return createLinks(*pHash, mp, app.list), nil
}
