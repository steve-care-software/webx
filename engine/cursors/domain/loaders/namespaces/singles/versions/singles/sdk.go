package singles

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/versions"
)

// Builder represents the single builder
type Builder interface {
	Create() Builder
	WithVersion(version versions.Version) Builder
	WithWorkspace() Builder
	WithMaster(master Single) Builder
	Now() (Single, error)
}

// Single represents the single namespace
type Single interface {
	Version() versions.Version
	HasWorkspace() bool
	Workspace()
	HasMaster() bool
	Master() Single
}
