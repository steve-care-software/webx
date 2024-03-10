package queries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/queries/conditions"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a query builder
type Builder interface {
	Create() Builder
	WithEntity(entity string) Builder
	WithCondition(condition conditions.Condition) Builder
	WithFields(fields []string) Builder
	Now() (Query, error)
}

// Query represents a query
type Query interface {
	Hash() hash.Hash
	Entity() string
	Condition() conditions.Condition
	HasFields() bool
	Fields() []string
}
