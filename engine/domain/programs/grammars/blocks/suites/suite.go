package suites

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/lexers"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/validations"
)

type suite struct {
	name        string
	input       []byte
	isFail      bool
	lexer       lexers.Lexer
	validations validations.Validations
}

func createSuite(
	name string,
	input []byte,
	isFail bool,
) Suite {
	return createSuiteInternally(
		name,
		input,
		isFail,
		nil,
		nil,
	)
}

func createSuiteWithLexer(
	name string,
	input []byte,
	isFail bool,
	lexer lexers.Lexer,
) Suite {
	return createSuiteInternally(
		name,
		input,
		isFail,
		lexer,
		nil,
	)
}

func createSuiteWithValidations(
	name string,
	input []byte,
	isFail bool,
	validations validations.Validations,
) Suite {
	return createSuiteInternally(
		name,
		input,
		isFail,
		nil,
		validations,
	)
}

func createSuiteWithLexerAndValidations(
	name string,
	input []byte,
	isFail bool,
	lexer lexers.Lexer,
	validations validations.Validations,
) Suite {
	return createSuiteInternally(
		name,
		input,
		isFail,
		lexer,
		validations,
	)
}

func createSuiteInternally(
	name string,
	input []byte,
	isFail bool,
	lexer lexers.Lexer,
	validations validations.Validations,
) Suite {
	out := suite{
		name:        name,
		input:       input,
		isFail:      isFail,
		lexer:       lexer,
		validations: validations,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Input returns the input
func (obj *suite) Input() []byte {
	return obj.input
}

// IsFail returns true if expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}

// HasLexer returns true if there is a lexer, false otherwise
func (obj *suite) HasLexer() bool {
	return obj.lexer != nil
}

// Lexer returns the lexer, if any
func (obj *suite) Lexer() lexers.Lexer {
	return obj.lexer
}

// HasValidations returns true if there is validations, false otherwise
func (obj *suite) HasValidations() bool {
	return obj.validations != nil
}

// Validations returns the validations, if any
func (obj *suite) Validations() validations.Validations {
	return obj.validations
}
