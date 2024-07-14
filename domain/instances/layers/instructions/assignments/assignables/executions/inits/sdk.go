package inits

import "github.com/steve-care-software/historydb/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an init builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	WithContext(context string) Builder
	Now() (Init, error)
}

// Init represents an init
type Init interface {
	Hash() hash.Hash
	Path() string
	Name() string
	Description() string
	Context() string
}
