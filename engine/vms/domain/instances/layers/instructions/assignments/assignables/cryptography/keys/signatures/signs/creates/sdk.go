package creates

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the create adapter
type Adapter interface {
	ToBytes(ins Create) ([]byte, error)
	ToInstance(bytes []byte) (Create, error)
}

// Builder represents a builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithPrivateKey(pk string) Builder
	Now() (Create, error)
}

// Create represents a create
type Create interface {
	Hash() hash.Hash
	Message() string
	PrivateKey() string
}
