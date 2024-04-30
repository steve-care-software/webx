package relationals

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isAnd       bool
	isOr        bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isAnd:       false,
		isOr:        false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// IsAnd adds an and to the builder
func (app *builder) IsAnd() Builder {
	app.isAnd = true
	return app
}

// IsOr adds an or to the builder
func (app *builder) IsOr() Builder {
	app.isOr = true
	return app
}

// Now builds a new Relational instance
func (app *builder) Now() (Relational, error) {
	isAnd := "false"
	isOr := "false"
	if app.isAnd {
		isAnd = "true"
	}

	if app.isOr {
		isOr = "true"
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(isAnd),
		[]byte(isOr),
	})

	if err != nil {
		return nil, err
	}

	if app.isAnd {
		return createRelationalWithAnd(*pHash), nil
	}

	if app.isOr {
		return createRelationalWithOr(*pHash), nil
	}

	return nil, errors.New("the Relational is invalid")
}
