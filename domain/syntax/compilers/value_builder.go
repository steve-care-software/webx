package compilers

import "github.com/steve-care-software/syntax/domain/syntax/criterias"

type valueBuilder struct {
	constant interface{}
	criteria criterias.Criteria
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		constant: nil,
		criteria: nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant interface{}) ValueBuilder {
	app.constant = constant
	return app
}

// WithCriteria adds a criteria to the builder
func (app *valueBuilder) WithCriteria(criteria criterias.Criteria) ValueBuilder {
	app.criteria = criteria
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.criteria != nil {
		return createValueWithCriteria(app.criteria), nil
	}

	if app.constant != nil {
		return createValueWithConstant(app.constant), nil
	}

	return createValueWithNil(), nil
}
