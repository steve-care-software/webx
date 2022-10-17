package elements

import (
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(hashAdapter)
}

// Builder represents an elements builder
type Builder interface {
	Create() Builder
	WithList(list []Element) Builder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	Hash() hash.Hash
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithName(name string) ElementBuilder
	WithCriteria(criteria criterias.Criteria) ElementBuilder
	Now() (Element, error)
}

// Element represents a route element
type Element interface {
	Hash() hash.Hash
	Name() string
	Criteria() criterias.Criteria
}
