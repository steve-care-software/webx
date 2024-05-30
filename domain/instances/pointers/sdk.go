package pointers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewPointerBuilder creates a new pointer builder
func NewPointerBuilder() PointerBuilder {
	hashAdapter := hash.NewAdapter()
	return createPointerBuilder(
		hashAdapter,
	)
}

// Adapter represents the pointers adapter
type Adapter interface {
	ToBytes(ins Pointers) ([]byte, error)
	ToInstance(bytes []byte) (Pointers, error)
}

// Builder represents the pointers builder
type Builder interface {
	Create() Builder
	WithList(list []Pointer) Builder
	Now() (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	Hash() hash.Hash
	List() []Pointer
	Match(executed [][]string) ([]Pointer, error)
}

// PointerBuilder represents the pointer builder
type PointerBuilder interface {
	Create() PointerBuilder
	WithPath(path []string) PointerBuilder
	WithLoader(loader conditions.Condition) PointerBuilder
	WithCanceller(canceller conditions.Condition) PointerBuilder
	IsActive() PointerBuilder
	Now() (Pointer, error)
}

// Resource represents a resource
type Pointer interface {
	Hash() hash.Hash
	Path() []string
	IsActive() bool
	HasLoader() bool
	Loader() conditions.Condition
	HasCanceller() bool
	Canceller() conditions.Condition
}
