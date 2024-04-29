package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new condition resource builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the resource adapter
type Adapter interface {
	ToBytes(ins Resource) ([]byte, error)
	ToInstance(bytes []byte) (Resource, error)
}

// Builder represents a condition resource builder
type Builder interface {
	Create() Builder
	WithCode(code uint) Builder
	IsRaisedInLayer() Builder
	Now() (Resource, error)
}

// Resource represents a condition resource
type Resource interface {
	Hash() hash.Hash
	Code() uint
	IsRaisedInLayer() bool
}
