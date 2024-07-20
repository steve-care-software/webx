package keys

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the key adapter
type Adapter interface {
	ToBytes(ins Key) ([]byte, error)
	ToInstance(bytes []byte) (Key, error)
}

// Builder represents a key builder
type Builder interface {
	Create() Builder
	WithEncryption(enc encryptions.Encryption) Builder
	WithSignature(sig signatures.Signature) Builder
	Now() (Key, error)
}

// Key represents a key
type Key interface {
	Hash() hash.Hash
	IsEncryption() bool
	Encryption() encryptions.Encryption
	IsSignature() bool
	Signature() signatures.Signature
}
