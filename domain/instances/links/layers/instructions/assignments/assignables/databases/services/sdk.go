package services

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the service adapter
type Adapter interface {
	ToBytes(ins Service) ([]byte, error)
	ToInstance(bytes []byte) (Service, error)
}

// Builder represents a service builder
type Builder interface {
	Create() Builder
	IsBegin() Builder
	Now() (Service, error)
}

// Service represents a service
type Service interface {
	Hash() hash.Hash
	IsBegin() bool
}
