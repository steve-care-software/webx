package validates

import (
	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the validate adapter
type Adapter interface {
	ToBytes(ins Validate) ([]byte, error)
	ToInstance(bytes []byte) (Validate, error)
}

// Builder represents a validate builder
type Builder interface {
	Create() Builder
	WithSignature(sig string) Builder
	WithMessage(msg string) Builder
	WithPublicKey(pubKey string) Builder
	Now() (Validate, error)
}

// Validate represents a validate
type Validate interface {
	Hash() hash.Hash
	Signature() string
	Message() string
	PublicKey() string
}
