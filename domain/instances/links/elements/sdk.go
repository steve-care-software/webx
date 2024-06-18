package elements

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
)

// NewBuilder creates a new elements builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementBuilder(
		hashAdapter,
	)
}

// Adapter represents the elements adapter
type Adapter interface {
	ToBytes(ins Elements) ([]byte, error)
	ToInstance(bytes []byte) (Elements, error)
}

// Builder represents elements builder
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

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithLayer(layer []string) ElementBuilder
	WithCondition(condition conditions.Condition) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Layer() []string
	HasCondition() bool
	Condition() conditions.Condition
}
