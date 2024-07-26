package metadatas

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

// NewBuilder creates a new metadata builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents a metadata adapter
type Adapter interface {
	ToBytes(ins MetaData) ([]byte, error)
	ToInstance(bytes []byte) (MetaData, error)
}

// Builder represents a metadata builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (MetaData, error)
}

// MetaData represents a database metadata
type MetaData interface {
	Hash() hash.Hash
	Path() []string
	Name() string
	Description() string
}
