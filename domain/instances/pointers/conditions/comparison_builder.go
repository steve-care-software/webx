package conditions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions/operators"
)

type comparisonBuilder struct {
	hashAdapter hash.Adapter
	operator    operators.Operator
	condition   Condition
}

func createComparisonBuilder(
	hashAdapter hash.Adapter,
) ComparisonBuilder {
	out := comparisonBuilder{
		hashAdapter: hashAdapter,
		operator:    nil,
		condition:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *comparisonBuilder) Create() ComparisonBuilder {
	return createComparisonBuilder(
		app.hashAdapter,
	)
}

// WithOperator adds an operator to the builder
func (app *comparisonBuilder) WithOperator(operator operators.Operator) ComparisonBuilder {
	app.operator = operator
	return app
}

// WithCondition adds a condition to the builder
func (app *comparisonBuilder) WithCondition(condition Condition) ComparisonBuilder {
	app.condition = condition
	return app
}

// Now builds a new Comparison instance
func (app *comparisonBuilder) Now() (Comparison, error) {
	if app.operator == nil {
		return nil, errors.New("the operator is mandatory in order to build a Comparison instance")
	}

	if app.condition == nil {
		return nil, errors.New("the condition is mandatory in order to build a Comparison instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.operator.Hash().Bytes(),
		app.condition.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createComparison(*pHash, app.operator, app.condition), nil
}
