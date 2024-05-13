package deletes

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	index       string
	length      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		index:       "",
		length:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index string) Builder {
	app.index = index
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length string) Builder {
	app.length = length
	return app
}

// Now builds a new Delete instance
func (app *builder) Now() (Delete, error) {
	if app.index == "" {
		return nil, errors.New("the index is mandatory in order to build a Delete instance")
	}

	if app.length == "" {
		return nil, errors.New("the length is mandatory in order to build a Delete instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.index),
		[]byte(app.length),
	})

	if err != nil {
		return nil, err
	}

	return createDelete(*pHash, app.index, app.length), nil
}
