package values

import "github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions/parameters/values/references"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the value builder
type Builder interface {
	Create() Builder
	WithReference(reference references.Reference) Builder
	WithBytes(bytes []byte) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	IsReference() bool
	Reference() references.Reference
	IsBytes() bool
	Bytes() []byte
}
