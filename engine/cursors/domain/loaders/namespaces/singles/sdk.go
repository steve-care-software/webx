package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/cursors/namespaces/versions"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/namespaces"
)

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithNamespace(namespace namespaces.Namespace) Builder
	WithVersion(version versions.Version) Builder
	Now() (Single, error)
}

// Single represents the single namespace
type Single interface {
	Namespace() namespaces.Namespace
	HasVersion() bool
	Version() versions.Version
}
