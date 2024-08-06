package branches

import (
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/delimiters"
	"github.com/steve-care-software/webx/engine/cursors/domain/storages/originals"
)

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
	WithOriginal(original originals.Original) BranchBuilder
	WithStates(states delimiters.Delimiter) BranchBuilder
	WithMetaData(metaData delimiters.Delimiter) BranchBuilder
	WithChildren(children Branches) BranchBuilder
	Now() (Branch, error)
}

// Branch represents a branch
type Branch interface {
	Original() originals.Original
	HasStates() bool
	States() delimiters.Delimiter
	HasMetaData() bool
	MetaData() delimiters.Delimiter
	HasChildren() bool
	Children() Branches
}
