package links

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type operatorBuilder struct {
	hashAdapter hash.Adapter
	isAnd       bool
	isOr        bool
	isXor       bool
}

func createOperatorBuilder(
	hashAdapter hash.Adapter,
) OperatorBuilder {
	out := operatorBuilder{
		hashAdapter: hashAdapter,
		isAnd:       false,
		isOr:        false,
		isXor:       false,
	}

	return &out
}

// Create initializes the builder
func (app *operatorBuilder) Create() OperatorBuilder {
	return createOperatorBuilder(
		app.hashAdapter,
	)
}

// IsAnd flags the builder as an and
func (app *operatorBuilder) IsAnd() OperatorBuilder {
	app.isAnd = true
	return app
}

// IsOr flags the builder as an or
func (app *operatorBuilder) IsOr() OperatorBuilder {
	app.isOr = true
	return app
}

// IsXor flags the builder as a xor
func (app *operatorBuilder) IsXor() OperatorBuilder {
	app.isXor = true
	return app
}

// Now builds a new Operator instance
func (app *operatorBuilder) Now() (Operator, error) {
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
