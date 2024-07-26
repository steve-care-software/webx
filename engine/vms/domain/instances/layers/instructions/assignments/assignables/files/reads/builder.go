package reads

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	identifier  string
	index       string
	length      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		identifier:  "",
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

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier string) Builder {
	app.identifier = identifier
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

// Now builds a new Read instance
func (app *builder) Now() (Read, error) {
	if app.identifier == "" {
		return nil, errors.New("the identifier is mandatory in order to build a Read instance")
	}

	data := [][]byte{
		[]byte(app.identifier),
	}

	if app.index != "" {
		data = append(data, []byte(app.index))
	}

	if app.length != "" {
		data = append(data, []byte(app.length))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.index != "" && app.length != "" {
		return createReadWithIndexAndLength(*pHash, app.identifier, app.index, app.length), nil
	}

	if app.index != "" {
		return createReadWithIndex(*pHash, app.identifier, app.index), nil
	}

	if app.length != "" {
		return createReadWithLength(*pHash, app.identifier, app.length), nil
	}

	return createRead(*pHash, app.identifier), nil
}
