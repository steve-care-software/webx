package compilers

import "github.com/steve-care-software/webx/selectors/domain/selectors"

type valueBuilder struct {
	constant string
	selector selectors.Selector
}

func createValueBuilder() ValueBuilder {
	out := valueBuilder{
		constant: "",
		selector: nil,
	}

	return &out
}

// Create initializes the builder
func (app *valueBuilder) Create() ValueBuilder {
	return createValueBuilder()
}

// WithConstant adds a constant to the builder
func (app *valueBuilder) WithConstant(constant string) ValueBuilder {
	app.constant = constant
	return app
}

// WithSelector adds a selector to the builder
func (app *valueBuilder) WithSelector(selector selectors.Selector) ValueBuilder {
	app.selector = selector
	return app
}

// Now builds a new Value instance
func (app *valueBuilder) Now() (Value, error) {
	if app.selector != nil {
		return createValueWithSelector(app.selector), nil
	}

	if app.constant != "" {
		return createValueWithConstant(app.constant), nil
	}

	return createValueWithNil(), nil
}
