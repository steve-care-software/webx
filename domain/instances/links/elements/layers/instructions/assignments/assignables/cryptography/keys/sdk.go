package keys

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures"
)

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
