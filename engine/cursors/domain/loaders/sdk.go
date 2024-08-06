package loaders

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces"

// Factory represents a loader factory
type Factory interface {
	Create() (Loader, error)
}

// Builder represents a loader builder
type Builder interface {
	Create() Builder
	WithNamespace(namespace namespaces.Namespace)
	Now() (Loader, error)
}

// Loader represents a loader
type Loader interface {
	HasIdentity() bool
	//Identity() identities.Identity
	HasNamespace() bool
	Namespace() namespaces.Namespace
	HasBlockchain() bool
	//Blockchain() blockchains.Blockchain
}
