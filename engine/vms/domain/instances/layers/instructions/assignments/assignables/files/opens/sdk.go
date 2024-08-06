package opens

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an open builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithPermission(permission string) Builder
	Now() (Open, error)
}

// Open represents an open file
type Open interface {
	Hash() hash.Hash
	Path() string
	Permission() string
}
