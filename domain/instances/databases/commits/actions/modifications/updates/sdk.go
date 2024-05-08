package updates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithDelete(del deletes.Delete) Builder
	WithBytes(bytes []byte) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Hash() hash.Hash
	Delete() deletes.Delete
	Bytes() []byte
}
