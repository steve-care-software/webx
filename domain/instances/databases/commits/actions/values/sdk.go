package values

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values/transforms"
)

// Adapter represents the value adapter
type Adapter interface {
	ToBytes(ins Value) ([]byte, error)
	ToInstance(data []byte) (Value, error)
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithInstance(instance instances.Instance) Builder
	WithTransform(transform transforms.Transform) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Hash() hash.Hash
	IsInstance() bool
	Instance() instances.Instance
	IsTransform() bool
	Transform() transforms.Transform
}
