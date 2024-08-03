package pointers

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers/delimiters"
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

// NewBuilder creates a new builder for tests
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

// Builder represents a pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	List() []Pointer
	Retrieve(hash hash.Hash) (Pointer, error)
	RetrieveAll(hashes []hash.Hash) ([]Pointer, error)
}

// PointerBuilder represents a pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithHash(hash hash.Hash) PointerBuilder
	WithDelimiter(delimiter delimiters.Delimiter) PointerBuilder
	Now() (Pointer, error)
}

// Pointer represents a pointer
type Pointer interface {
	Hash() hash.Hash
	Delimiter() delimiters.Delimiter
}
