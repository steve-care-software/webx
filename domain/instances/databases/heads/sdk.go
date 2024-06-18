package heads

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the heads adapter
type Adapter interface {
	ToBytes(ins Head) ([]byte, error)
	ToInstance(bytes []byte) (Head, error)
}

// Builder represents an head builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithDescription(description string) Builder
	IsActive() Builder
	Now() (Head, error)
}

// Head represents a database head
type Head interface {
	Hash() hash.Hash
	Path() []string
	Description() string
	IsActive() bool
}
