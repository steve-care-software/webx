package versions

import "github.com/steve-care-software/webx/engine/bytes/domain/delimiters"

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
	WithName(name string) VersionBuilder
	WithDescription(description string) VersionBuilder
	IsDeleted() VersionBuilder
	WithIterations(iterations delimiters.Delimiter) VersionBuilder
	WithPrevious(previous Version) VersionBuilder
	Now() (Version, error)
}

// Version represents a version
type Version interface {
	Name() string
	Description() string
	IsDeleted() bool
	HasIterations() bool
	Iterations() delimiters.Delimiter
	HasPrevious() bool
	Previous() Version
}
