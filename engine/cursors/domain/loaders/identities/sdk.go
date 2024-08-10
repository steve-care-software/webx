package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys"
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
	WithAuthenticated(authenticated keys.Keys) Builder
	WithCurrent(current keys.Key) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	All() storages.Storages
	HasAuthenticated() bool
	Authenticated() keys.Keys
	HasCurrent() bool
	Current() keys.Key
}
