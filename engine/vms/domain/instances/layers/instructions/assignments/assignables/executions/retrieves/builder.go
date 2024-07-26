package retrieves

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	context     string
	index       string
	length      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		context:     "",
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

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
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

// Now builds a new Retrieve instance
func (app *builder) Now() (Retrieve, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build a Retrieve instance")
	}

	if app.index == "" {
		return nil, errors.New("the index is mandatory in order to build a Retrieve instance")
	}

	bytes := [][]byte{
		[]byte(app.context),
		[]byte(app.index),
	}

	if app.length != "" {
		bytes = append(bytes, []byte(app.length))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.length != "" {
		return createRetrieveWithLength(*pHash, app.index, app.context, app.length), nil
	}

	return createRetrieve(*pHash, app.index, app.context), nil
}
