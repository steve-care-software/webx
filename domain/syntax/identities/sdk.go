package identities

import (
	"time"

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

// Repository represents an identity repository
type Repository interface {
	List() ([]string, error)
	Retrieve(name string, password string) (Identity, error)
}

// Service represents an identity service
type Service interface {
	Insert(identity Identity, password string) error
	Update(identity Identity, currentPassword string, newPassword string) error
}
