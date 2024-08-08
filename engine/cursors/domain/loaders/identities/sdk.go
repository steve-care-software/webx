package identities

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/singles"
)

// Builder represents the identity builder
type Builder interface {
	Create() Builder
	WithAll(all []string) Builder
	WithCurrent(current singles.Single) Builder
	Now() (Identity, error)
}

// Identity represents an identity
type Identity interface {
	All() []string
	HasCurrent() bool
	Current() singles.Single
}
