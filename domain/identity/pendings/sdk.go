package pendings

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/syntax/domain/identity/cryptography/signatures"
)

// Builder represents a pendings builder
type Builder interface {
	Create() Builder
	WithList(list []Pending) Builder
	Now() (Pendings, error)
}

// Pendings represents a pendings
type Pendings interface {
	List() []Pending
}

// PendingBuilder represents a pending builder
type PendingBuilder interface {
	Create() PendingBuilder
	WithID(id uuid.UUID) PendingBuilder
	WithRing(ring []signatures.PublicKey) PendingBuilder
	WithPrivateKey(pk signatures.PrivateKey) PendingBuilder
	Now() (Pending, error)
}

// Pending represents a pending
type Pending interface {
	Public() Public
	PrivateKey() signatures.PrivateKey
}

// Public represents a public pending
type Public interface {
	ID() uuid.UUID
	Ring() []signatures.PublicKey
}

// Repository represents a pending repository
type Repository interface {
	List() []uuid.UUID
	Retrieve(id uuid.UUID) (Pending, error)
}

// Service represents a pending service
type Service interface {
	Save(pending Pending) error
	Delete(pending Pending) error
}
