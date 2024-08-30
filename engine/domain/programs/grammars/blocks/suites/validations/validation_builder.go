package validations

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"
)

type validationBuilder struct {
	variable variables.Variable
	isFail   bool
}

func createValidationBuilder() ValidationBuilder {
	out := validationBuilder{
		variable: nil,
		isFail:   false,
	}

	return &out
}

// Create initializes the builder
func (app *validationBuilder) Create() ValidationBuilder {
	return createValidationBuilder()
}

// WithVariable adds a variable to the builder
func (app *validationBuilder) WithVariable(variable variables.Variable) ValidationBuilder {
	app.variable = variable
	return app
}

// IsFail flags the builder as fail
func (app *validationBuilder) IsFail() ValidationBuilder {
	app.isFail = true
	return app
}

// Now builds a new Validation instance
func (app *validationBuilder) Now() (Validation, error) {
	if app.variable == nil {
		return nil, errors.New("the variable is mandatory in order to build a Validation instance")
	}

	return createValidation(
		app.variable,
		app.isFail,
	), nil
}
