package constants

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new constant builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a constant builder
type Builder interface {
	Create() Builder
	WithBool(boolValue bool) Builder
	WithBytes(bytes []byte) Builder
	Now() (Constant, error)
}

// Constant represents a constant assignable
type Constant interface {
	Hash() hash.Hash
	IsBool() bool
	Bool() *bool
	IsBytes() bool
	Bytes() []byte
}
