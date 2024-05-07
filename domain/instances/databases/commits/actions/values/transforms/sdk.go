package transforms

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the transform adapter
type Adapter interface {
	ToBytes(ins Transform) ([]byte, error)
	ToInstance(data []byte) (Transform, error)
}

// Builder represents a transform builder
type Builder interface {
	Create() Builder
	WithQuery(query []byte) Builder
	WithBytes(bytes []byte) Builder
	Now() (Transform, error)
}

// Transform represents a transform
type Transform interface {
	Hash() hash.Hash
	Query() []byte
	Bytes() []byte
}
