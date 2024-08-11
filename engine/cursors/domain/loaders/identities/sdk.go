package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers"
)

// NewBuilder creates a new builder for tests
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the identity builder
type Builder interface {
	Create() Builder
	WithAll(all storages.Storages) Builder
	WithAuthenticated(authenticated switchers.Switchers) Builder
	WithCurrent(current switchers.Switcher) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	All() storages.Storages
	HasAuthenticated() bool
	Authenticated() switchers.Switchers
	HasCurrent() bool
	Current() switchers.Switcher
}
