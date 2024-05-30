package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/references"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewLinkBuilder creates a new link builder instance
func NewLinkBuilder() LinkBuilder {
	hashAdapter := hash.NewAdapter()
	return createLinkBuilder(
		hashAdapter,
	)
}

// Adapter represents the links adapter
type Adapter interface {
	ToBytes(ins Links) ([]byte, error)
	ToInstance(bytes []byte) (Links, error)
}

// Builder represents the links builder
type Builder interface {
	Create() Builder
	WithList(list []Link) Builder
	Now() (Links, error)
}

// Links represents links
type Links interface {
	Hash() hash.Hash
	List() []Link
	Fetch(hash hash.Hash) (Link, error)
}

// Repository represents the links repository
type Repository interface {
	Retrieve(path []string) (Links, error)
}

// LinkAdapter represents the link adapter
type LinkAdapter interface {
	ToBytes(ins Link) ([]byte, error)
	ToInstance(bytes []byte) (Link, error)
}

// LinkBuilder represents a link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithElements(elements elements.Elements) LinkBuilder
	WithReferences(references references.References) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Elements() elements.Elements
	HasReferences() bool
	References() references.References
}

// LinkRepository represents the link repository
type LinkRepository interface {
	Retrieve() (Link, error)
	RetrieveFromPath(path hash.Hash) (Link, error)
}
