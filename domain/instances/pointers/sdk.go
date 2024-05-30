package pointers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
)

// Adapter represents the pointers adapter
type Adapter interface {
	ToBytes(ins Pointers) ([]byte, error)
	ToInstance(bytes []byte) (Pointers, error)
}

// Pointers represents pointers
type Pointers interface {
	Hash() hash.Hash
	List() []Pointer
	Match(executed [][]string) ([]Pointer, error)
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
