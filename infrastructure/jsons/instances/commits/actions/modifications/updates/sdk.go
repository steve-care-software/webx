package updates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithDelete(delete string) Builder
	WithBytes(bytes string) Builder
	Now() (Update, error)
}

// Update represents an update
type Update interface {
	Hash() hash.Hash
	Delete() string
	Bytes() string
}
