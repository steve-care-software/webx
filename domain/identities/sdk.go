package identities

import (
	"time"

	"github.com/steve-care-software/webx/domain/cryptography/encryptions/keys"
	"github.com/steve-care-software/webx/domain/cryptography/signatures"
	database_identities "github.com/steve-care-software/webx/domain/databases/identities"
	"github.com/steve-care-software/webx/domain/identities/modifications"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an identity adapter
type Adapter interface {
	ToDatabase(ins Identity) (database_identities.Identity, error)
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
