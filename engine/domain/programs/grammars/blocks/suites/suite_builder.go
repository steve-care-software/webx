package suites

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/lexers"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/validations"
)

type suiteBuilder struct {
	name        string
	input       []byte
	isFail      bool
	lexer       lexers.Lexer
	validations validations.Validations
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		name:        "",
		input:       nil,
		isFail:      false,
		lexer:       nil,
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

// WithInput adds a input to the builder
func (app *suiteBuilder) WithInput(input []byte) SuiteBuilder {
	app.input = input
	return app
}

// IsFail flags the suite as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// WithLexer adds a lexer to the builder
func (app *suiteBuilder) WithLexer(lexer lexers.Lexer) SuiteBuilder {
	app.lexer = lexer
	return app
}

// WithValidations add the validations to the builder
func (app *suiteBuilder) WithValidations(validations validations.Validations) SuiteBuilder {
	app.validations = validations
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build a Suite instance")
	}

	if app.validations != nil {
		return createSuiteWithValidations(app.name, app.input, app.isFail, app.validations), nil
	}

	return createSuite(app.name, app.input, app.isFail), nil
}
