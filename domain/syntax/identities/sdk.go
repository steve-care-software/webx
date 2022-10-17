package identities

import (
	"time"

	"github.com/steve-care-software/syntax/domain/syntax/databases"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/encryptions/keys"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/signatures"
	"github.com/steve-care-software/syntax/domain/syntax/identities/modifications"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithSignature(signature signatures.PrivateKey) Builder
	WithEncryption(encryption keys.PrivateKey) Builder
	WithDatabases(databases databases.Databases) Builder
	WithModifications(modifications modifications.Modifications) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Name() string
	Signature() signatures.PrivateKey
	Encryption() keys.PrivateKey
	CreatedOn() time.Time
	HasDatabases() bool
	Databases() databases.Databases
	HasModifications() bool
	Modifications() modifications.Modifications
}
