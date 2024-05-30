package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/links/elements/conditions/resources"
)

// NewBuilder creates a new condition builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the condition adapter
type Adapter interface {
	ToBytes(ins Condition) ([]byte, error)
	ToInstance(bytes []byte) (Condition, error)
}

// Builder represents condition builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource) Builder
	WithNext(next Condition) Builder
	Now() (Condition, error)
}

// Condition represents a condition
type Condition interface {
	Hash() hash.Hash
	Resource() resources.Resource
	HasNext() bool
	Next() Condition
}
