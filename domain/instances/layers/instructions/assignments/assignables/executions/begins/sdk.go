package begins

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a begin builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithContext(context string) Builder
	Now() (Begin, error)
}

// Begin represents a begin
type Begin interface {
	Hash() hash.Hash
	Path() string
	Context() string
}
