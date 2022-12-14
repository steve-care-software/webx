package identities

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithModifications(modifications modifications.Modifications) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Signature() signatures.PrivateKey
	Encryption() keys.PrivateKey
	CreatedOn() time.Time
	Modifications() modifications.Modifications
}
