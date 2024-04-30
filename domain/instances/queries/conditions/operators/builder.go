package operators

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/integers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators/relationals"
)

type builder struct {
	hashAdapter hash.Adapter
	isEqual     bool
	relational  relationals.Relational
	integer     integers.Integer
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		isEqual:     false,
		relational:  nil,
		integer:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithRelational adds a relational to the builder
func (app *builder) WithRelational(relational relationals.Relational) Builder {
	app.relational = relational
	return app
}

// WithInteger adds an integer to the builder
func (app *builder) WithInteger(integer integers.Integer) Builder {
	app.integer = integer
	return app
}

// IsEqual flags the builder as equal
func (app *builder) IsEqual() Builder {
	app.isEqual = true
	return app
}

// Now builds a new Operator instance
func (app *builder) Now() (Operator, error) {
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
