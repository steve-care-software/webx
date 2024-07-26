package inits

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the init adapter
type Adapter interface {
	ToBytes(ins Init) ([]byte, error)
	ToInstance(bytes []byte) (Init, error)
}

// Builder represents an init builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (Init, error)
}

// Init represents an init
type Init interface {
	Hash() hash.Hash
	Path() string
	Name() string
	Description() string
}
