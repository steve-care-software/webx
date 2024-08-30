package validations

import "github.com/steve-care-software/webx/engine/domain/stacks/frames/variables"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewValidationBuilder creates a new validation builder
func NewValidationBuilder() ValidationBuilder {
	return createValidationBuilder()
}

// Builder represents validations builder
type Builder interface {
	Create() Builder
	WithList(list []Validation) Builder
	Now() (Validations, error)
}

// Validations represents validations
type Validations interface {
	List() []Validation
}

// ValidationBuilder represents a validation builder
type ValidationBuilder interface {
	Create() ValidationBuilder
	WithVariable(variable variables.Variable) ValidationBuilder
	IsFail() ValidationBuilder
	Now() (Validation, error)
}

// Validation represents a validation
type Validation interface {
	Variable() variables.Variable
	IsFail() bool
}
