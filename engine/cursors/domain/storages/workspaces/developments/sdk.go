package developments

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/loaders/resources/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

// Adapter represents a development adapter
type Adapter interface {
	InstancesToBytes(ins Developments) ([]byte, error)
	BytesToInstances(data []byte) (Developments, []byte, error)
	InstanceToBytes(ins Development) ([]byte, error)
	BytesToInstance(data []byte) (Development, []byte, error)
}

// Builder represents a developments builder
type Builder interface {
	Create() Builder
	WithList(list []Development) Builder
	Now() (Developments, error)
}

// Developments represents a developments instance
type Developments interface {
	List() []Development
}

// DevelopmentBuilder represents a development builder
type DevelopmentBuilder interface {
	Create() DevelopmentBuilder
	WithOriginal(original originals.Original) DevelopmentBuilder
	WithBranches(branches delimiters.Delimiter) DevelopmentBuilder
	Now() (Development, error)
}

// Development represents a development iteration
type Development interface {
	Original() originals.Original
	HasBranch() bool
	Branch() delimiters.Delimiter
}
