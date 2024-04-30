package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/operators"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/queries/conditions/resources"
)

// NewBuilder creates a new builder instance
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

// Builder represents a condition builder
type Builder interface {
	Create() Builder
	WithPointer(pointer pointers.Pointer) Builder
	WithOperator(operator operators.Operator) Builder
	WithElement(element Element) Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Pointer() pointers.Pointer
	Operator() operators.Operator
	Element() Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithCondition(condition Condition) ElementBuilder
	WithResource(resource resources.Resource) ElementBuilder
	Now() (Element, error)
}

// Element represents a conditional element
type Element interface {
	Hash() hash.Hash
	IsCondition() bool
	Condition() Condition
	IsResource() bool
	Resource() resources.Resource
}
