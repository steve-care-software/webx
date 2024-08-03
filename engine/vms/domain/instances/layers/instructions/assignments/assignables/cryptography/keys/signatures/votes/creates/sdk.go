package creates

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// NewBuilder creates a new builder
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

// Builder creates a create builder
type Builder interface {
	Create() Builder
	WithMessage(message string) Builder
	WithRing(ring string) Builder
	WithPrivateKey(pk string) Builder
	Now() (Create, error)
}

// Create represents a cretae vote
type Create interface {
	Hash() hash.Hash
	Message() string
	Ring() string
	PrivateKey() string
}
