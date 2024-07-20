package deletes

import (
	"errors"

	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        string
	index       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        "",
		index:       "",
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
func (app *builder) WithList(list string) Builder {
	app.list = list
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index string) Builder {
	app.index = index
	return app
}

// Now builds a new Delete instance
func (app *builder) Now() (Delete, error) {
	if app.list == "" {
		return nil, errors.New("the list is mandatory in order to build a Delete instance")
	}

	if app.index == "" {
		return nil, errors.New("the index is mandatory in order to build a Delete instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.list),
		[]byte(app.index),
	})

	if err != nil {
		return nil, err
	}

	return createDelete(*pHash, app.list, app.index), nil
}
