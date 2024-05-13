package fetches

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a fetch builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithIndex(index string) Builder
	Now() (Fetch, error)
}

// Fetch represents a fetch
type Fetch interface {
	Hash() hash.Hash
	List() string
	Index() string
}
