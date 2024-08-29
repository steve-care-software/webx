package references

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the reference builder
type Builder interface {
	Create() Builder
	WithElement(element elements.Element) Builder
	WithIndex(index uint) Builder
	Now() (Reference, error)
}

// Reference represents a reference
type Reference interface {
	Element() elements.Element
	Index() uint
}
