package resources

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages"
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/switchers"
)

// Builder represents the resource builder
type Builder interface {
	Create() Builder
	WithAll(all storages.Storages) Builder
	WithLoaded(loaded switchers.Switchers) Builder
	WithCurrent(current switchers.Switcher) Builder
	Now() (Resource, error)
}

// Resource represents an resource
type Resource interface {
	All() storages.Storages
	HasLoaded() bool
	Loaded() switchers.Switchers
	HasCurrent() bool
	Current() switchers.Switcher
}
