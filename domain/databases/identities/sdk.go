package identities

import "github.com/steve-care-software/webx/domain/databases/entities"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an identity adapter
type Adapter interface {
	ToContent(ins Identity) ([]byte, error)
}

// Builder represents an identity builder
type Builder interface {
	Create() Builder
	WithIdentifier(identifier entities.Identifier) Builder
	WithModifications(modifications entities.Identifiers) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	Identifier() entities.Identifier
	Modifications() entities.Identifiers
}
