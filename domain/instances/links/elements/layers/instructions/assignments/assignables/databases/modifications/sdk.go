package modifications

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

// Builder represents a modification builder
type Builder interface {
	Create() Builder
	WithInsert(insert string) Builder
	WithDelete(delete string) Builder
	Now() (Modification, error)
}

// Modification represents a modification
type Modification interface {
	Hash() hash.Hash
	IsInsert() bool
	Insert() string
	IsDelete() bool
	Delete() string
}
