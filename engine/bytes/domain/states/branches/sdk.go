package branches

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/delimiters"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers"
)

// NewBranchBuilder creates a new branch builder
func NewBranchBuilder() BranchBuilder {
	return createBranchBuilder()
}

// Adapter represents a branch adapter
type Adapter interface {
	InstancesToBytes(ins Branches) ([]byte, error)
	BytesToInstances(data []byte) (Branches, []byte, error)
	InstanceToBytes(ins Branch) ([]byte, error)
	BytesToInstance(data []byte) (Branch, []byte, error)
}

// Builder represents a branches builder
type Builder interface {
	Create() Builder
	WithList(list []Branch) Builder
	Now() (Branches, error)
}

// Branches represents branches
type Branches interface {
	List() []Branch
}

// BranchBuilder represents a branch builder
type BranchBuilder interface {
	Create() BranchBuilder
	WithName(name string) BranchBuilder
	WithLayers(layers layers.Layers) BranchBuilder
	WithMetaData(metaData delimiters.Delimiter) BranchBuilder
	WithChildren(children Branches) BranchBuilder
	IsDeleted() BranchBuilder
	Now() (Branch, error)
}

// Branch represents a branch
type Branch interface {
	Name() string
	IsDeleted() bool
	HasLayers() bool
	Layers() layers.Layers
	HasMetaData() bool
	MetaData() delimiters.Delimiter
	HasChildren() bool
	Children() Branches
}
