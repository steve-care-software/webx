package fetches

import "github.com/steve-care-software/datastencil/domain/hash"

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
