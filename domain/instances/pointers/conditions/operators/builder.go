package operators

import (
	"errors"

	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	isAnd       bool
	isOr        bool
	isXor       bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isAnd:       false,
		isOr:        false,
		isXor:       false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// IsAnd flags the builder as an and
func (app *builder) IsAnd() Builder {
	app.isAnd = true
	return app
}

// IsOr flags the builder as an or
func (app *builder) IsOr() Builder {
	app.isOr = true
	return app
}

// IsXor flags the builder as a xor
func (app *builder) IsXor() Builder {
	app.isXor = true
	return app
}

// Now builds a new Operator instance
func (app *builder) Now() (Operator, error) {
	data := [][]byte{}
	if app.isAnd {
		data = append(data, []byte("isAnd"))
	}

	if app.isOr {
		data = append(data, []byte("isOr"))
	}

	if app.isXor {
		data = append(data, []byte("isXor"))
	}

	if len(data) != 1 {
		return nil, errors.New("the operator is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isAnd {
		return createOperatorWithIsAnd(*pHash), nil
	}

	if app.isOr {
		return createOperatorWithIsOr(*pHash), nil
	}

	return createOperatorWithIsXor(*pHash), nil
}
