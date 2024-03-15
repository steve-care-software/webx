package pointers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        []string
	identifier  hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        nil,
		identifier:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier hash.Hash) Builder {
	app.identifier = identifier
	return app
}

// Now builds a new pointer instance
func (app *builder) Now() (Pointer, error) {
	if app.path != nil && len(app.path) <= 0 {
		app.path = nil
	}

	if app.path == nil {
		return nil, errors.New("the path is mandatory in order to build a Pointer instance")
	}

	if app.identifier == nil {
		return nil, errors.New("the identifier is mandatory in order to build a Pointer instance")
	}

	data := [][]byte{
		app.identifier.Bytes(),
	}

	for _, onePath := range app.path {
		data = append(data, []byte(onePath))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createPointer(*pHash, app.path, app.identifier), nil
}
