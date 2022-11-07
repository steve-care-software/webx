package identities

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithEntity(entity entities.Entity) Builder
	WithModifications(modifications entities.Identifiers) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Entity() entities.Entity
	Modifications() entities.Identifiers
}
