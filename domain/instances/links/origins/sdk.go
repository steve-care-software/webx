package origins

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

// NewBuilder creates a new origin builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewValueBuilder creates a new origin value builder
func NewValueBuilder() ValueBuilder {
	hashAdapter := hash.NewAdapter()
	return createValueBuilder(
		hashAdapter,
	)
}

// Adapter represents the origin adapter
type Adapter interface {
	ToBytes(ins Origin) ([]byte, error)
	ToInstance(bytes []byte) (Origin, error)
}

// Builder represents the origin builder
type Builder interface {
	Create() Builder
	WithResource(resource resources.Resource) Builder
	WithOperator(operator operators.Operator) Builder
	WithNext(next Value) Builder
	Now() (Origin, error)
}

// Origin represents an origin
type Origin interface {
	Hash() hash.Hash
	Resource() resources.Resource
	Operator() operators.Operator
	Next() Value
}

// ValueBuilder represents the value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithResource(resource resources.Resource) ValueBuilder
	WithOrigin(origin Origin) ValueBuilder
	Now() (Value, error)
}

// Value represents an origin value
type Value interface {
	Hash() hash.Hash
	IsResource() bool
	Resource() resources.Resource
	IsOrigin() bool
	Origin() Origin
}
