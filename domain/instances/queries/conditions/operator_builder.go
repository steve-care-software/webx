package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type operatorBuilder struct {
	hashAdapter hash.Adapter
	isEqual     bool
	relational  RelationalOperator
	integer     IntegerOperator
}

func createOperatorBuilder(
	hashAdapter hash.Adapter,
) OperatorBuilder {
	out := operatorBuilder{
		hashAdapter: hashAdapter,
		isEqual:     false,
		relational:  nil,
		integer:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *operatorBuilder) Create() OperatorBuilder {
	return createOperatorBuilder(
		app.hashAdapter,
	)
}

// WithRelational adds a relational to the builder
func (app *operatorBuilder) WithRelational(relational RelationalOperator) OperatorBuilder {
	app.relational = relational
	return app
}

// WithInteger adds an integer to the builder
func (app *operatorBuilder) WithInteger(integer IntegerOperator) OperatorBuilder {
	app.integer = integer
	return app
}

// IsEqual flags the builder as equal
func (app *operatorBuilder) IsEqual() OperatorBuilder {
	app.isEqual = true
	return app
}

// Now builds a new Operator instance
func (app *operatorBuilder) Now() (Operator, error) {
	data := [][]byte{}
	if app.isEqual {
		data = append(data, []byte("isEqual"))
	}

	if app.relational != nil {
		data = append(data, app.relational.Hash().Bytes())
	}

	if app.integer != nil {
		data = append(data, app.integer.Hash().Bytes())
	}

	if len(data) != 1 {
		return nil, errors.New("the Operator is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.isEqual {
		return createOperatorWithEqual(*pHash), nil
	}

	if app.relational != nil {
		return createOperatorWithRelational(*pHash, app.relational), nil
	}

	return createOperatorWithInteger(*pHash, app.integer), nil
}
