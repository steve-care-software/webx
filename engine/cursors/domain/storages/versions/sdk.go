package versions

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Adapter represents a version adapter
type Adapter interface {
	InstancesToBytes(ins Versions) ([]byte, error)
	BytesToInstances(data []byte) (Versions, []byte, error)
	InstanceToBytes(ins Version) ([]byte, error)
	BytesToInstance(data []byte) (Version, []byte, error)
}

// Builder represents a versions builder
type Builder interface {
	Create() Builder
	WithList(list []Version) Builder
	Now() (Versions, error)
}

// Versions represents versions
type Versions interface {
	List() []Version
}

// VersionBuilder represents a version builder
type VersionBuilder interface {
	Create() VersionBuilder
	WithOriginal(original originals.Original) VersionBuilder
	WithWorkspaces(workspaces delimiters.Delimiter) VersionBuilder
	WithMaster(master delimiters.Delimiter) VersionBuilder
	Now() (Version, error)
}

// Version represents a version
type Version interface {
	Original() originals.Original
	HasWorkspaces() bool
	Workspaces() delimiters.Delimiter
	HasMaster() bool
	Master() delimiters.Delimiter
}
