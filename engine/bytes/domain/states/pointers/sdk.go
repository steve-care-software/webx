package pointers

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	return createPointerBuilder()
}

// Adapter represents a pointers adapter
type Adapter interface {
	InstancesToBytes(ins Pointers) ([]byte, error)
	BytesToInstances(data []byte) (Pointers, []byte, error)
	InstanceToBytes(ins Pointer) ([]byte, error)
	BytesToInstance(data []byte) (Pointer, []byte, error)
}

// Builder represents the pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
	NextIndex() uint64
	Fetch(delimiter delimiters.Delimiter) (Pointer, error)
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithDelimiter(delimiter delimiters.Delimiter) PointerBuilder
	IsDeleted() PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Delimiter() delimiters.Delimiter
	IsDeleted() bool
}
