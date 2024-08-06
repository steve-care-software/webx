package keynames

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// Now builds a new Keyname instance
func (app *builder) Now() (Keyname, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Keyname instance")
	}

	pHash, err := app.hashAdapter.FromBytes([]byte(app.name))
	if err != nil {
		return nil, err
	}

	return createKeyname(
		*pHash,
		app.name,
	), nil
}
