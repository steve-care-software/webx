package listers

import "github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the lister builder
type Builder interface {
	Create() Builder
	WithKeyname(keyname string) Builder
	WithRetrieval(retrieval retrievals.Retrieval) Builder
	Now() (Lister, error)
}

// Lister represents a lister
type Lister interface {
	Keyname() string
	Retrieval() retrievals.Retrieval
}
