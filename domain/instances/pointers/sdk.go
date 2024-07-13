package pointers

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/conditions"
	"github.com/steve-care-software/historydb/domain/hash"
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
	First() Pointer
	Match(executed [][]string) []Pointer
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

// Resource represents a pointer
type Pointer interface {
	Hash() hash.Hash
	Path() []string
	IsActive() bool
	HasLoader() bool
	Loader() conditions.Condition
	HasCanceller() bool
	Canceller() conditions.Condition
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithBasePath(basePath []string) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a pointer repository
type Repository interface {
	Retrieve(path []string) (Pointers, error)
	Match(path []string, history [][]string) (Pointers, error)
	Fetch(path []string, history [][]string) ([]byte, error)
}
