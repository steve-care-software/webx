package validates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Builder represents a validate builder
type Builder interface {
	Create() Builder
	WithSignature(sig string) Builder
	WithMessage(msg Builder) Builder
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
