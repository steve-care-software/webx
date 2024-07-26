package files

import "github.com/steve-care-software/webx/engine/databases/entities/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a file builder
type Builder interface {
	Create() Builder
	WithClose(close string) Builder
	WithDelete(del string) Builder
	Now() (File, error)
}

// File represents a file
type File interface {
	Hash() hash.Hash
	IsClose() bool
	Close() string
	IsDelete() bool
	Delete() string
}
