package values

import (
	"github.com/steve-care-software/webx/engine/domain/programs/instructions/tokens/elements/syscalls/values/parameters"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Builder represents the values builder
type Builder interface {
	Create() Builder
	WithList(list []Value) Builder
	Now() (Values, error)
}

// Values represents values
type Values interface {
	List() []Value
}

// ValueBuilder represents the value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithParameter(parameter parameters.Parameter) ValueBuilder
	WithToken(token string) ValueBuilder
	Now() (Value, error)
}

// Value represents a syscall value
type Value interface {
	IsParameter() bool
	Parameter() parameters.Parameter
	IsToken() bool
	Token() string
}
