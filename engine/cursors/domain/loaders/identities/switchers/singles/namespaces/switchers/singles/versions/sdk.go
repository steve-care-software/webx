package versions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/switchers/singles/namespaces/switchers/singles/versions/switchers"
)

// Builder represents a versions builder
type Builder interface {
	Create() Builder
	WithList(list []Version) Builder
	Now() (Versions, error)
}

// Versions represents versions
type Versions interface {
	List() []Version
	Names() []string
	Fetch(name string) (Version, error)
}

// VersionBuilder represents a version builder
type VersionBuilder interface {
	Create() VersionBuilder
	WithAll(storages storages.Storages) VersionBuilder
	WithLoaded(loaded switchers.Switchers) VersionBuilder
	WithCurrent(current switchers.Switcher) VersionBuilder
	Now() (Version, error)
}

// Version represents a version
type Version interface {
	All() storages.Storages
	HasLoaded() bool
	Loaded() switchers.Switchers
	HasCurrent() bool
	Current() switchers.Switcher
}
