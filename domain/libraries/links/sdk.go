package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/links/origins"
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

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	hashAdapter := hash.NewAdapter()
	return createElementsBuilder(
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

// NewConditionBuilder creates a new condition builder
func NewConditionBuilder() ConditionBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionBuilder(
		hashAdapter,
	)
}

// NewConditionValueBuilder creates a new condition value builder
func NewConditionValueBuilder() ConditionValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionValueBuilder(
		hashAdapter,
	)
}

// NewConditionResourceBuilder creates a new condition resource builder
func NewConditionResourceBuilder() ConditionResourceBuilder {
	hashAdapter := hash.NewAdapter()
	return createConditionResourceBuilder(
		hashAdapter,
	)
}

// Adapter represents the links adapter
type Adapter interface {
	ToData(ins Links) ([]byte, error)
	ToInstance(data []byte) (Links, error)
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
	FetchByExecutedLayers(layerHashes []hash.Hash) (Link, error)
}

// Repository represents the links repository
type Repository interface {
	Retrieve(path []string) (Links, error)
}

// LinkAdapter represents the link adapter
type LinkAdapter interface {
	ToData(ins Link) ([]byte, error)
	ToInstance(data []byte) (Link, error)
}

// LinkBuilder represents a link builder
type LinkBuilder interface {
	Create() LinkBuilder
	WithOrigin(origin origins.Origin) LinkBuilder
	WithElements(elements Elements) LinkBuilder
	Now() (Link, error)
}

// Link represents a link
type Link interface {
	Hash() hash.Hash
	Origin() origins.Origin
	Elements() Elements
}

// LinkRepository represents the link repository
type LinkRepository interface {
	Retrieve(path []string) (Link, error)
}

// ElementsBuilder represents elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
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
	WithLayer(layer hash.Hash) ElementBuilder
	WithLayerBytes(layerBytes []byte) ElementBuilder
	WithCondition(condition Condition) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Layer() hash.Hash
	HasCondition() bool
	Condition() Condition
}

// ConditionBuilder represents condition builder
type ConditionBuilder interface {
	Create() ConditionBuilder
	WithResource(resource ConditionResource) ConditionBuilder
	WithNext(next ConditionValue) ConditionBuilder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Resource() ConditionResource
	HasNext() bool
	Next() ConditionValue
}

// ConditionValueBuilder represents a condition value builder
type ConditionValueBuilder interface {
	Create() ConditionValueBuilder
	WithResource(resource ConditionResource) ConditionValueBuilder
	WithCondition(condition Condition) ConditionValueBuilder
	Now() (ConditionValue, error)
}

// ConditionValue represents a condition value
type ConditionValue interface {
	Hash() hash.Hash
	IsResource() bool
	Resource() ConditionResource
	IsCondition() bool
	Condition() Condition
}

// ConditionResourceBuilder represents a condition resource builder
type ConditionResourceBuilder interface {
	Create() ConditionResourceBuilder
	WithCode(code uint) ConditionResourceBuilder
	IsRaisedInLayer() ConditionResourceBuilder
	Now() (ConditionResource, error)
}

// ConditionResource represents a condition resource
type ConditionResource interface {
	Hash() hash.Hash
	Code() uint
	IsRaisedInLayer() bool
}
