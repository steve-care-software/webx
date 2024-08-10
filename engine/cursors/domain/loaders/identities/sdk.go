package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/storages"
)

// NewBuilder creates a new builder for tests
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the identity builder
type Builder interface {
	Create() Builder
	WithAll(all storages.Storages) Builder
	WithAuthenticated(authenticated singles.Singles) Builder
	WithCurrent(current singles.Single) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	All() storages.Storages
	HasAuthenticated() bool
	Authenticated() singles.Singles
	HasCurrent() bool
	Current() singles.Single
}
