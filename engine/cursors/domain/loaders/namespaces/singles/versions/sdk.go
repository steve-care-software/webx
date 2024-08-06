package versions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/namespaces/singles/versions/singles"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/versions"
)

// Builder represents a version builder
type Builder interface {
	Create() Builder
	WithAll(all versions.Versions) Builder
	WithCurrent(current singles.Single) Builder
	Now() (Version, error)
}

// Version represents a version
type Version interface {
	All() versions.Versions
	HasCurrent() bool
	Current() singles.Single
}
