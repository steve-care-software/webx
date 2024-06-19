package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/references"
)

// NewBuilder creates a new link builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the link adapter
type Adapter interface {
	ToBytes(ins Link) ([]byte, error)
	ToInstance(bytes []byte) (Link, error)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithElements(elements elements.Elements) Builder
	WithReferences(references references.References) Builder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Elements() elements.Elements
	HasReferences() bool
	References() references.References
}

// RepositoryBuilder represents a repository builder
type RepositoryBuilder interface {
	Create() RepositoryBuilder
	WithBasePath(basePath []string) RepositoryBuilder
	Now() (Repository, error)
}

// Repository represents a link repository
type Repository interface {
	Retrieve(path []string, history [][]string) (Link, error)
}
