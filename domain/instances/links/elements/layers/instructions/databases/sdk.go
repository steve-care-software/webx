package databases

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

// Adapter represents the database adapter
type Adapter interface {
	ToBytes(ins Database) ([]byte, error)
	ToInstance(bytes []byte) (Database, error)
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithSave(save string) Builder
	WithDelete(delete string) Builder
	Now() (Database, error)
}

// Database represents a database instruction
type Database interface {
	Hash() hash.Hash
	IsSave() bool
	Save() string
	IsDelete() bool
	Delete() string
}
