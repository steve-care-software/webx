package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type relationalOperatorBuilder struct {
	hashAdapter hash.Adapter
	isAnd       bool
	isOr        bool
}

func createRelationalOperatorBuilder(
	hashAdapter hash.Adapter,
) RelationalOperatorBuilder {
	out := relationalOperatorBuilder{
		hashAdapter: hashAdapter,
		isAnd:       false,
		isOr:        false,
	}

	return &out
}

// Create initializes the builder
func (app *relationalOperatorBuilder) Create() RelationalOperatorBuilder {
	return createRelationalOperatorBuilder(
		app.hashAdapter,
	)
}

// IsAnd adds an and to the builder
func (app *relationalOperatorBuilder) IsAnd() RelationalOperatorBuilder {
	app.isAnd = true
	return app
}

// IsOr adds an or to the builder
func (app *relationalOperatorBuilder) IsOr() RelationalOperatorBuilder {
	app.isOr = true
	return app
}

// Now builds a new RelationalOperator instance
func (app *relationalOperatorBuilder) Now() (RelationalOperator, error) {
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
		return createRelationalOperatorWithAnd(*pHash), nil
	}

	if app.isOr {
		return createRelationalOperatorWithOr(*pHash), nil
	}

	return nil, errors.New("the RelationalOperator is invalid")
}
