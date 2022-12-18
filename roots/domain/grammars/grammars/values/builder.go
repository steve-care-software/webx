package values

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	name        string
	pNumber     *byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		name:        "",
		pNumber:     nil,
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

// WithNumber adds a number to the builder
func (app *builder) WithNumber(number byte) Builder {
	app.pNumber = &number
	return app
}

// Now builds a new Value instance
func (app *builder) Now() (Value, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Value instance")
	}

	if app.pNumber == nil {
		return nil, errors.New("the value is mandatory in order to build a Value instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.name),
		[]byte{
			*app.pNumber,
		},
	})

	if err != nil {
		return nil, err
	}

	return createValue(*pHash, app.name, *app.pNumber), nil
}
