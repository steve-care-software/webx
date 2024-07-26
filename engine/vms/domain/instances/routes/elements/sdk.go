package elements

import (
	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
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
	InstancesToBytes(ins Elements) ([]byte, error)
	BytesToInstances(bytes []byte) (Elements, error)
	InstanceToBytes(ins Element) ([]byte, error)
	BytesToInstance(bytes []byte) (Element, error)
}

// Builder represents the elements builder
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
	WithLayer(layer hash.Hash) ElementBuilder
	WithBytes(bytes []byte) ElementBuilder
	WithString(str string) ElementBuilder
	Now() (Element, error)
}

// Element represents a route element
type Element interface {
	Hash() hash.Hash
	IsLayer() bool
	Layer() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsString() bool
	String() string
}
