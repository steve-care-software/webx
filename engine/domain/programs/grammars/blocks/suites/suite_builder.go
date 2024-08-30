package suites

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/validations"
)

type suiteBuilder struct {
	name        string
	value       []byte
	isFail      bool
	validations validations.Validations
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		name:        "",
		value:       nil,
		isFail:      false,
		validations: nil,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
}

// WithName adds a name to the builder
func (app *suiteBuilder) WithName(name string) SuiteBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *suiteBuilder) WithValue(value []byte) SuiteBuilder {
	app.value = value
	return app
}

// IsFail flags the suite as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// WithValidations add the validations to the builder
func (app *suiteBuilder) WithValidations(validations validations.Validations) SuiteBuilder {
	app.validations = validations
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.value != nil && len(app.value) <= 0 {
		app.value = nil
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a Suite instance")
	}

	if app.validations != nil {
		return createSuiteWithValidations(app.name, app.value, app.isFail, app.validations), nil
	}

	return createSuite(app.name, app.value, app.isFail), nil
}
