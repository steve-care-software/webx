package locations

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents location builder
type Builder interface {
	Create() Builder
	WithSingle(single []byte) Builder
	WithList(list [][]byte) Builder
	Now() (Location, error)
}

// Location represents the layer location
type Location interface {
	Hash() hash.Hash
	IsSingle() bool
	Single() []byte
	IsList() bool
	List() [][]byte
}
