package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewLogicBuilder creates a new logic builder
func NewLogicBuilder() LogicBuilder {
	hashAdapter := hash.NewAdapter()
	return createLogicBuilder(
		hashAdapter,
	)
}

// Builder represents the logics builder
type Builder interface {
	Create() Builder
	WithList(list []Logic) Builder
	Now() (Logics, error)
}

// Logics represents logics
type Logics interface {
	Hash() hash.Hash
	List() []Logic
}

// LogicBuilder represents a logic builder
type LogicBuilder interface {
	Create() LogicBuilder
	WithLink(link links.Link) LogicBuilder
	WithLayers(layers layers.Layers) LogicBuilder
	Now() (Logic, error)
}

// Logic represents a logic
type Logic interface {
	Hash() hash.Hash
	Link() links.Link
	Layers() layers.Layers
}
