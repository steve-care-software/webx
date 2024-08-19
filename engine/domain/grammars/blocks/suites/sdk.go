package suites

import "github.com/steve-care-software/webx/engine/domain/grammars/tokens/elements"

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
	WithElement(element elements.Element) SuiteBuilder
	IsFail() SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Name() string
	Element() elements.Element
	IsFail() bool
}
