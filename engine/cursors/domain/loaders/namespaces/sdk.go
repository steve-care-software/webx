package namespaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/headers/namespaces"
)

// Builder represents a namespace builder
type Builder interface {
	Create() Builder
	WithAll(all namespaces.Namespaces) Builder
	WithCurrent(current singles.Single) Builder
	Now() (Namespace, error)
}

// Namespace represents a namespace
type Namespace interface {
	All() namespaces.Namespaces
	Current() singles.Single
}
