package suites

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/lexers"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/suites/validations"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	return createSuiteBuilder()
}

// Builder represents the suites builder
type Builder interface {
	Create() Builder
	WithList(list []Suite) Builder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// SuiteBuilder represents the suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithName(name string) SuiteBuilder
	WithInput(input []byte) SuiteBuilder
	WithLexer(lexer lexers.Lexer) SuiteBuilder
	WithValidations(validations validations.Validations) SuiteBuilder
	IsFail() SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Name() string
	Input() []byte
	IsFail() bool
	HasLexer() bool
	Lexer() lexers.Lexer
	HasValidations() bool
	Validations() validations.Validations
}
