package namespaces

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers"
)

// Builder represents a namespaces builder
type Builder interface {
	Create() Builder
	WithList(list []Namespace) Builder
	Now() (Namespaces, error)
}

// Namespaces represents namespaces
type Namespaces interface {
	List() []Namespace
	Names() []string
	Fetch(name string) (Namespace, error)
}

// NamespaceBuilder represents a namespace builder
type NamespaceBuilder interface {
	Create() NamespaceBuilder
	WithAll(storages storages.Storages) NamespaceBuilder
	WithLoaded(loaded switchers.Switchers) NamespaceBuilder
	WithCurrent(current switchers.Switcher) NamespaceBuilder
	Now() (Namespace, error)
}

// Namespace represents a namespace
type Namespace interface {
	All() storages.Storages
	HasLoaded() bool
	Loaded() switchers.Switchers
	HasCurrent() bool
	Current() switchers.Switcher
}
