package values

import (
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

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
	WithElement(element elements.Element) ValueBuilder
	Now() (Value, error)
}

// Value represents a syscall value
type Value interface {
	IsParameter() bool
	Parameter() parameters.Parameter
	IsElement() bool
	Element() elements.Element
}
