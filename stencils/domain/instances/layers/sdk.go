package layers

import (
	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/outputs"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/references"
)

// NewBuilder creates a new layer builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the layer adapter
type Adapter interface {
	ToBytes(ins Layer) ([]byte, error)
	ToInstance(bytes []byte) (Layer, error)
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions instructions.Instructions) Builder
	WithOutput(output outputs.Output) Builder
	WithInput(input string) Builder
	WithReferences(references references.References) Builder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	Hash() hash.Hash
	Instructions() instructions.Instructions
	Output() outputs.Output
	Input() string
	HasReferences() bool
	References() references.References
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithBasePath(basePath []string) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a layer repository
type Repository interface {
	Retrieve(path []string, history [][]string) (Layer, error)
}
