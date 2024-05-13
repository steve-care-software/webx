package inserts

import "github.com/steve-care-software/datastencil/domain/hash"

// Builder represents an insert builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithElement(element string) Builder
	Now() (Insert, error)
}

// Insert represents an insert
type Insert interface {
	Hash() hash.Hash
	List() string
	Element() string
}
