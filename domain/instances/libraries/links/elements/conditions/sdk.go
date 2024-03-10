package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/elements/conditions/resources"
)

// NewBuilder creates a new condition builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
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

// Builder represents condition builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource) Builder
	WithNext(next ConditionValue) Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Resource() resources.Resource
	HasNext() bool
	Next() ConditionValue
}

// ConditionValueBuilder represents a condition value builder
type ConditionValueBuilder interface {
	Create() ConditionValueBuilder
	WithResource(resource resources.Resource) ConditionValueBuilder
	WithCondition(condition Condition) ConditionValueBuilder
	Now() (ConditionValue, error)
}

// ConditionValue represents a condition value
type ConditionValue interface {
	Hash() hash.Hash
	IsResource() bool
	Resource() resources.Resource
	IsCondition() bool
	Condition() Condition
}
