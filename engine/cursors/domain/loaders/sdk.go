package loaders

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"
)

// Factory represents a loader factory
type Factory interface {
	Create() (Loader, error)
}

// Builder represents a loader builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	WithInitialLoader(loader Loader) Builder
	Now() (Loader, error)
}

// Loader represents a loader
type Loader interface {
	HasIdentity() bool
	Identity() identities.Identity
	HasNamespace() bool
}
