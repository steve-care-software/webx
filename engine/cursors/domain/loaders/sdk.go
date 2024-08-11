package loaders

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities"
)

// NewFactory creates a new factory instance
func NewFactory() Factory {
	builder := NewBuilder()
	return createFactory(
		builder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Factory represents a loader factory
type Factory interface {
	Create() (Loader, error)
}

// Builder represents a loader builder
type Builder interface {
	Create() Builder
	WithIdentity(identity identities.Identity) Builder
	Now() (Loader, error)
}

// Loader represents a loader
type Loader interface {
	HasIdentity() bool
	Identity() identities.Identity
}
