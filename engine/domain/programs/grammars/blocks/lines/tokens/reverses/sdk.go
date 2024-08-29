package reverses

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the reverse builder
type Builder interface {
	Create() Builder
	WithEscape(escape elements.Element) Builder
	Now() (Reverse, error)
}

// Reverse represents the reverse builder
type Reverse interface {
	HasEscape() bool
	Escape() elements.Element
}
